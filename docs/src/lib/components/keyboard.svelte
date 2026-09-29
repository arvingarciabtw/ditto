<script lang="ts">
	import { onMount } from "svelte";
	import { SvelteSet } from "svelte/reactivity";

	const labelByCode: Record<string, string> = {
		Escape: "Esc",
		F1: "F1",
		F2: "F2",
		F3: "F3",
		F4: "F4",
		F5: "F5",
		F6: "F6",
		F7: "F7",
		F8: "F8",
		F9: "F9",
		F10: "F10",
		F11: "F11",
		F12: "F12",
		PrintScreen: "Prt",
		ScrollLock: "Slk",
		Pause: "Fls",
		Backquote: "`",
		Digit1: "1",
		Digit2: "2",
		Digit3: "3",
		Digit4: "4",
		Digit5: "5",
		Digit6: "6",
		Digit7: "7",
		Digit8: "8",
		Digit9: "9",
		Digit0: "0",
		Minus: "-",
		Equal: "=",
		Backspace: "<--",
		Insert: "Ins",
		Home: "Hme",
		PageUp: "PgU",
		Tab: "Tab",
		KeyQ: "Q",
		KeyW: "W",
		KeyE: "E",
		KeyR: "R",
		KeyT: "T",
		KeyY: "Y",
		KeyU: "U",
		KeyI: "I",
		KeyO: "O",
		KeyP: "P",
		BracketLeft: "[",
		BracketRight: "]",
		Backslash: "\\",
		Delete: "Del",
		End: "End",
		PageDown: "PgD",
		CapsLock: "Caps",
		KeyA: "A",
		KeyS: "S",
		KeyD: "D",
		KeyF: "F",
		KeyG: "G",
		KeyH: "H",
		KeyJ: "J",
		KeyK: "K",
		KeyL: "L",
		Semicolon: ";",
		Quote: "'",
		Enter: "Enter",
		ShiftLeft: "Shift",
		ShiftRight: "Shift",
		KeyZ: "Z",
		KeyX: "X",
		KeyC: "C",
		KeyV: "V",
		KeyB: "B",
		KeyN: "N",
		KeyM: "M",
		Comma: ",",
		Period: ".",
		Slash: "/",
		ControlLeft: "Ctrl",
		ControlRight: "Ctrl",
		MetaLeft: "⌘",
		MetaRight: "⌘",
		AltLeft: "Alt",
		AltRight: "Alt",
		Space: "Space",
		Fn: "Fn",
		ArrowLeft: "←",
		ArrowDown: "↓",
		ArrowUp: "↑",
		ArrowRight: "→",
		NumLock: "Nlk",
		NumpadDivide: "/",
		NumpadMultiply: "*",
		NumpadSubtract: "-",
		NumpadAdd: "+",
		Numpad7: "7",
		Numpad8: "8",
		Numpad9: "9",
		Numpad4: "4",
		Numpad5: "5",
		Numpad6: "6",
		Numpad1: "1",
		Numpad2: "2",
		Numpad3: "3",
		NumpadEnter: "E",
		Numpad0: "0",
		NumpadDecimal: "."
	};

	const shiftedLabelByCode: Record<string, string> = {
		Backquote: "~",
		Digit1: "!",
		Digit2: "@",
		Digit3: "#",
		Digit4: "$",
		Digit5: "%",
		Digit6: "^",
		Digit7: "&",
		Digit8: "*",
		Digit9: "(",
		Digit0: ")",
		Minus: "_",
		Equal: "+",
		BracketLeft: "{",
		BracketRight: "}",
		Backslash: "|",
		Semicolon: ":",
		Quote: '"',
		Comma: "<",
		Period: ">",
		Slash: "?"
	};

	const latchModifierByCode: Record<string, string> = {
		CapsLock: "CapsLock",
		NumLock: "NumLock",
		ScrollLock: "ScrollLock"
	};

	const pressedCodes = new SvelteSet<string>();
	const latchedCodes = new SvelteSet<string>();

	function isShiftDown() {
		return pressedCodes.has("ShiftLeft") || pressedCodes.has("ShiftRight");
	}

	function isLit(code: string) {
		return pressedCodes.has(code) || latchedCodes.has(code);
	}

	function isCapsOn() {
		return latchedCodes.has("CapsLock");
	}

	function labelOf(code: string) {
		const label = labelByCode[code] ?? "";
		if (!isShiftDown()) return label;
		// Caps + Shift inverts letter case, like a real board. "Key*" is only KeyA..KeyZ.
		if (isCapsOn() && code.startsWith("Key")) return label.toLowerCase();
		return shiftedLabelByCode[code] ?? label;
	}

	function flipLatch(code: string) {
		if (latchedCodes.has(code)) latchedCodes.delete(code);
		else latchedCodes.add(code);
	}

	function syncLatches(event: KeyboardEvent) {
		for (const [code, modifier] of Object.entries(latchModifierByCode)) {
			if (event.getModifierState(modifier)) latchedCodes.add(code);
			else latchedCodes.delete(code);
		}
	}

	function resyncLatches() {
		const probe = new KeyboardEvent("keydown");
		for (const [code, modifier] of Object.entries(latchModifierByCode)) {
			// Only ever turn a latch on. Browsers disagree on whether getModifierState()
			// reflects real hardware for synthetic events (Firefox does, Chrome always
			// says false), and a false result must never clear a correct indicator.
			if (probe.getModifierState(modifier)) latchedCodes.add(code);
		}
	}

	onMount(() => {
		const press = (event: KeyboardEvent) => {
			if (event.code in latchModifierByCode) {
				// A lock key physically toggles, but its own keydown/keyup can still report
				// the previous state (and macOS never fires keyup for it), so trust the press
				// instead of the reported state. `repeat` guards against auto-repeat flicker.
				if (!event.repeat) flipLatch(event.code);
				return;
			}
			syncLatches(event);
			pressedCodes.add(event.code);
		};
		const release = (event: KeyboardEvent) => {
			if (event.code in latchModifierByCode) return;
			syncLatches(event);
			pressedCodes.delete(event.code);
		};
		const clear = () => pressedCodes.clear();

		window.addEventListener("keydown", press);
		window.addEventListener("keyup", release);
		window.addEventListener("blur", clear);
		window.addEventListener("focus", resyncLatches);
		document.addEventListener("visibilitychange", resyncLatches);

		return () => {
			window.removeEventListener("keydown", press);
			window.removeEventListener("keyup", release);
			window.removeEventListener("blur", clear);
			window.removeEventListener("focus", resyncLatches);
			document.removeEventListener("visibilitychange", resyncLatches);
		};
	});
</script>

{#snippet gridkey(
	code: string,
	finger: string,
	topBorder: boolean = false,
	rightBorder: boolean = true,
	bottomBorder: boolean = true,
	leftBorder: boolean = false
)}
	<div
		class={`key ${isLit(code) ? "active" : "inactive"} ${finger} ${topBorder ? "top-border" : "no-top-border"} ${rightBorder ? "right-border" : "no-right-border"} ${bottomBorder ? "bottom-border" : "no-bottom-border"} ${leftBorder ? "left-border" : "no-left-border"}`}
	>
		<p>{labelOf(code)}</p>
	</div>
{/snippet}

{#snippet abskey(code: string, finger: string)}
	<div class={`key abs ${isLit(code) ? "active" : "inactive"} ${finger}`}>
		<p>{labelOf(code)}</p>
	</div>
{/snippet}

<div class="grid-keyboard">
	<div class="row top">
		{@render gridkey("Escape", "pinky")}
		{@render gridkey("", "")}
		{@render gridkey("F1", "pinky")}
		{@render gridkey("F2", "ring")}
		{@render gridkey("F3", "middle")}
		{@render gridkey("F4", "index")}
		{@render gridkey("", "")}
		{@render gridkey("F5", "index")}
		{@render gridkey("F6", "index")}
		{@render gridkey("F7", "index")}
		{@render gridkey("F8", "middle")}
		{@render gridkey("", "")}
		{@render gridkey("F9", "ring")}
		{@render gridkey("F10", "pinky")}
		{@render gridkey("F11", "pinky")}
		{@render gridkey("F12", "pinky")}
		{@render gridkey("", "", false, true, false)}
		{@render gridkey("PrintScreen", "index")}
		{@render gridkey("ScrollLock", "middle")}
		{@render gridkey("Pause", "ring")}
		{@render gridkey("", "", false, false, false)}
		{@render gridkey("", "", false, false)}
	</div>
	<div class="row number">
		{@render gridkey("Backquote", "pinky")}
		{@render gridkey("Digit1", "pinky")}
		{@render gridkey("Digit2", "ring")}
		{@render gridkey("Digit3", "middle")}
		{@render gridkey("Digit4", "index")}
		{@render gridkey("Digit5", "index")}
		{@render gridkey("Digit6", "index")}
		{@render gridkey("Digit7", "index")}
		{@render gridkey("Digit8", "middle")}
		{@render gridkey("Digit9", "ring")}
		{@render gridkey("Digit0", "pinky")}
		{@render gridkey("Minus", "pinky")}
		{@render gridkey("Equal", "pinky")}
		{@render gridkey("Backspace", "pinky")}
		{@render gridkey("", "", false, true, false)}
		{@render gridkey("Insert", "index")}
		{@render gridkey("Home", "middle")}
		{@render gridkey("PageUp", "ring")}
		{@render gridkey("", "", false, true, false)}
		{@render gridkey("NumLock", "index")}
		{@render gridkey("NumpadDivide", "middle")}
		{@render gridkey("NumpadMultiply", "ring")}
		{@render gridkey("NumpadSubtract", "pinky", false, false)}
	</div>
	<div class="row qwer">
		{@render gridkey("Tab", "pinky")}
		{@render gridkey("KeyQ", "pinky")}
		{@render gridkey("KeyW", "ring")}
		{@render gridkey("KeyE", "middle")}
		{@render gridkey("KeyR", "index")}
		{@render gridkey("KeyT", "index")}
		{@render gridkey("KeyY", "index")}
		{@render gridkey("KeyU", "index")}
		{@render gridkey("KeyI", "middle")}
		{@render gridkey("KeyO", "ring")}
		{@render gridkey("KeyP", "pinky")}
		{@render gridkey("BracketLeft", "pinky")}
		{@render gridkey("BracketRight", "pinky")}
		{@render gridkey("Backslash", "pinky")}
		{@render gridkey("", "", false, true, false)}
		{@render gridkey("Delete", "index")}
		{@render gridkey("End", "middle")}
		{@render gridkey("PageDown", "ring")}
		{@render gridkey("", "", false, true, false)}
		{@render gridkey("Numpad7", "index")}
		{@render gridkey("Numpad8", "middle")}
		{@render gridkey("Numpad9", "ring")}
		{@render abskey("NumpadAdd", "ring")}
		{@render gridkey("", "", false, false, false)}
	</div>
	<div class="row asdf">
		{@render gridkey("CapsLock", "pinky")}
		{@render gridkey("KeyA", "pinky")}
		{@render gridkey("KeyS", "ring")}
		{@render gridkey("KeyD", "middle")}
		{@render gridkey("KeyF", "index")}
		{@render gridkey("KeyG", "index")}
		{@render gridkey("KeyH", "index")}
		{@render gridkey("KeyJ", "index")}
		{@render gridkey("KeyK", "middle")}
		{@render gridkey("KeyL", "ring")}
		{@render gridkey("Semicolon", "pinky")}
		{@render gridkey("Quote", "pinky")}
		{@render gridkey("Enter", "pinky")}
		{@render gridkey("", "", false, true, false)}
		{@render gridkey("Numpad4", "index")}
		{@render gridkey("Numpad5", "middle")}
		{@render gridkey("Numpad6", "ring")}
		{@render gridkey("", "", false, false)}
	</div>
	<div class="row zxcv">
		{@render gridkey("ShiftLeft", "pinky")}
		{@render gridkey("KeyZ", "pinky")}
		{@render gridkey("KeyX", "ring")}
		{@render gridkey("KeyC", "middle")}
		{@render gridkey("KeyV", "index")}
		{@render gridkey("KeyB", "index")}
		{@render gridkey("KeyN", "index")}
		{@render gridkey("KeyM", "index")}
		{@render gridkey("Comma", "middle")}
		{@render gridkey("Period", "ring")}
		{@render gridkey("Slash", "pinky")}
		{@render gridkey("ShiftRight", "pinky")}
		{@render gridkey("", "", false, false, false)}
		{@render gridkey("", "", false, true, true, false)}
		{@render gridkey("ArrowUp", "middle", true)}
		{@render gridkey("", "", false, false, true, false)}
		{@render gridkey("", "", false, true, false)}
		{@render gridkey("Numpad1", "index")}
		{@render gridkey("Numpad2", "middle")}
		{@render gridkey("Numpad3", "ring")}
		{@render abskey("NumpadEnter", "ring")}
		{@render gridkey("", "", false, false, false)}
	</div>
	<div class="row bottom">
		{@render gridkey("ControlLeft", "pinky")}
		{@render gridkey("MetaLeft", "ring")}
		{@render gridkey("AltLeft", "ring")}
		{@render gridkey("Space", "thumb")}
		{@render gridkey("AltRight", "ring")}
		{@render gridkey("MetaRight", "ring")}
		{@render gridkey("Fn", "ring")}
		{@render gridkey("ControlRight", "pinky")}
		{@render gridkey("", "")}
		{@render gridkey("ArrowLeft", "index")}
		{@render gridkey("ArrowDown", "middle")}
		{@render gridkey("ArrowRight", "ring")}
		{@render gridkey("", "")}
		{@render gridkey("Numpad0", "index")}
		{@render gridkey("NumpadDecimal", "ring")}
		{@render gridkey("", "", false, false)}
	</div>
</div>

<blockquote>
	this keyboard updates, try typing! take note that the actual tui program visualizes a bit
	differently due to terminal limitations (⌐■-■)
</blockquote>

<style>
	blockquote {
		padding-left: 1rem;
		margin-top: 2rem;
		border-left: 4px solid var(--surface1);
		color: var(--overlay1);
		line-height: 2;
	}

	.grid-keyboard {
		width: 100%;
		max-width: 820px;
		border: 1px solid var(--surface2);
		border-radius: 4px;
		--square: 34px;
		.row {
			position: relative;
			display: grid;
			align-items: center;
		}
		.row.top {
			grid-template-columns:
				repeat(6, var(--square)) 26px repeat(4, var(--square)) 26px repeat(4, var(--square))
				26px repeat(3, var(--square)) 26px 1fr;
		}
		.row.number {
			grid-template-columns:
				repeat(13, var(--square)) 86px 26px repeat(3, var(--square))
				26px repeat(4, var(--square));
		}
		.row.qwer {
			grid-template-columns:
				calc(2 * var(--square)) repeat(12, var(--square)) 52px 26px repeat(3, var(--square))
				26px repeat(4, var(--square));
		}
		.row.asdf {
			grid-template-columns: 60px repeat(11, var(--square)) 94px 154px repeat(4, var(--square));
		}
		.row.zxcv {
			grid-template-columns:
				100px repeat(10, var(--square)) 88px 26px 34px var(--square)
				34px 26px repeat(4, var(--square));
		}
		.row.bottom {
			grid-template-columns:
				60px var(--square) 51px 202px 51px repeat(2, var(--square))
				62px 26px repeat(3, var(--square)) 26px 68px repeat(2, var(--square));
			.key {
				border-bottom: none;
			}
		}
		.key {
			padding: 0 4px;
			height: 38px;
			display: grid;
			align-items: center;
			justify-items: center;
			border-right: 1px solid var(--surface2);
			border-bottom: 1px solid var(--surface2);
		}
		.key.abs {
			position: absolute;
			top: 20px;
			left: 793px;
			border: none;
			color: var(--pink);
		}
		.key.top-border {
			border-top: 1px solid var(--surface2);
		}
		.key.no-top-border {
			border-top: none;
		}
		.key.right-border {
			border-right: 1px solid var(--surface2);
		}
		.key.no-right-border {
			border-right: none;
		}
		.key.bottom-border {
			border-bottom: 1px solid var(--surface2);
		}
		.key.no-bottom-border {
			border-bottom: none;
		}
		.key.left-border {
			border-left: 1px solid var(--surface2);
		}
		.key.no-left-border {
			border-left: none;
		}
		.key.pinky {
			--finger: var(--pink);
			color: var(--finger);
		}
		.key.ring {
			--finger: var(--sapphire);
			color: var(--finger);
		}
		.key.middle {
			--finger: var(--green);
			color: var(--finger);
		}
		.key.index {
			--finger: var(--peach);
			color: var(--finger);
		}
		.key.thumb {
			--finger: var(--red);
			color: var(--finger);
		}
		.key.active {
			color: var(--crust);
			background: var(--finger);
			border-color: var(--finger);
		}
	}

	@media (max-width: 900px) {
		.keyboard-wrapper {
			font-size: 12px;
		}
		.grid-keyboard {
			max-width: 694px;
			font-size: 12px;
			--square: 28px;
			.row.top {
				grid-template-columns:
					repeat(6, var(--square)) 26px repeat(4, var(--square)) 26px repeat(4, var(--square))
					26px repeat(3, var(--square)) 26px 1fr;
			}
			.row.number {
				grid-template-columns:
					repeat(13, var(--square)) 80px 26px repeat(3, var(--square))
					26px repeat(4, var(--square));
			}
			.row.qwer {
				grid-template-columns:
					calc(2 * var(--square)) repeat(12, var(--square)) 52px 26px repeat(3, var(--square))
					26px repeat(4, var(--square));
			}
			.row.asdf {
				grid-template-columns: 60px repeat(11, var(--square)) 76px 136px repeat(4, var(--square));
			}
			.row.zxcv {
				grid-template-columns:
					100px repeat(10, var(--square)) 64px 26px 28px var(--square)
					28px 26px repeat(4, var(--square));
			}
			.row.bottom {
				grid-template-columns:
					50px var(--square) 45px 170px 45px repeat(2, var(--square))
					50px 26px repeat(3, var(--square)) 26px 56px repeat(2, var(--square));
				.key {
					border-bottom: none;
				}
			}
			.key {
				padding: 0 2px;
				height: 36px;
			}
			.key.abs {
				top: 18px;
				left: 672px;
			}
		}
	}
	@media (max-width: 750px) {
		.keyboard-wrapper {
			font-size: 10px;
		}
		.grid-keyboard {
			max-width: 564px;
			font-size: 10px;
			--square: 23px;
			.row.top {
				grid-template-columns:
					repeat(6, var(--square)) 20px repeat(4, var(--square)) 20px repeat(4, var(--square))
					20px repeat(3, var(--square)) 20px 1fr;
			}
			.row.number {
				grid-template-columns:
					repeat(13, var(--square)) 63px 20px repeat(3, var(--square))
					20px repeat(4, var(--square));
			}
			.row.qwer {
				grid-template-columns:
					calc(2 * var(--square)) repeat(12, var(--square)) 40px 20px repeat(3, var(--square))
					20px repeat(4, var(--square));
			}
			.row.asdf {
				grid-template-columns: 40px repeat(11, var(--square)) 69px 109px repeat(4, var(--square));
			}
			.row.zxcv {
				grid-template-columns:
					70px repeat(10, var(--square)) 62px 20px 23px var(--square)
					23px 20px repeat(4, var(--square));
			}
			.row.bottom {
				grid-template-columns:
					40px var(--square) 36px 141px 36px repeat(2, var(--square))
					40px 20px repeat(3, var(--square)) 20px 46px repeat(2, var(--square));
				.key {
					border-bottom: none;
				}
			}
			.key {
				padding: 0 2px;
				height: 28px;
			}
			.key.abs {
				top: 14px;
				left: 546px;
			}
		}
	}
</style>
