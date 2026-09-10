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

	const pressedCodes = new SvelteSet<string>();

	function isPressed(label: string) {
		for (const code of pressedCodes) {
			if (labelByCode[code] === label) return true;
		}
		return false;
	}

	onMount(() => {
		const press = (event: KeyboardEvent) => {
			if (!labelByCode[event.code]) return;
			pressedCodes.add(event.code);
		};
		const release = (event: KeyboardEvent) => {
			pressedCodes.delete(event.code);
		};
		const clear = () => pressedCodes.clear();

		window.addEventListener("keydown", press);
		window.addEventListener("keyup", release);
		window.addEventListener("blur", clear);

		return () => {
			window.removeEventListener("keydown", press);
			window.removeEventListener("keyup", release);
			window.removeEventListener("blur", clear);
		};
	});
</script>

{#snippet gridkey(
	content: string,
	finger: string,
	topBorder: boolean = false,
	rightBorder: boolean = true,
	bottomBorder: boolean = true,
	leftBorder: boolean = false
)}
	<div
		class={`key ${finger} ${topBorder ? "top-border" : "no-top-border"} ${rightBorder ? "right-border" : "no-right-border"} ${bottomBorder ? "bottom-border" : "no-bottom-border"} ${leftBorder ? "left-border" : "no-left-border"}`}
	>
		<p class:active={isPressed(content)}>{content}</p>
	</div>
{/snippet}

<div class="grid-keyboard">
	<div class="row top">
		{@render gridkey("Esc", "pinky")}
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
		{@render gridkey("Prt", "index")}
		{@render gridkey("Slk", "middle")}
		{@render gridkey("Fls", "ring")}
		{@render gridkey("", "", false, false, false)}
		{@render gridkey("", "", false, false)}
	</div>
	<div class="row number">
		{@render gridkey("`", "pinky")}
		{@render gridkey("1", "pinky")}
		{@render gridkey("2", "ring")}
		{@render gridkey("3", "middle")}
		{@render gridkey("4", "index")}
		{@render gridkey("5", "index")}
		{@render gridkey("6", "index")}
		{@render gridkey("7", "index")}
		{@render gridkey("8", "middle")}
		{@render gridkey("9", "ring")}
		{@render gridkey("0", "pinky")}
		{@render gridkey("-", "pinky")}
		{@render gridkey("=", "pinky")}
		{@render gridkey("<--", "pinky")}
		{@render gridkey("", "", false, true, false)}
		{@render gridkey("Ins", "index")}
		{@render gridkey("Hme", "middle")}
		{@render gridkey("PgU", "ring")}
		{@render gridkey("", "", false, true, false)}
		{@render gridkey("Nlk", "index")}
		{@render gridkey("/", "middle")}
		{@render gridkey("*", "ring")}
		{@render gridkey("-", "pinky", false, false)}
	</div>
	<div class="row qwer">
		{@render gridkey("Tab", "pinky")}
		{@render gridkey("Q", "pinky")}
		{@render gridkey("W", "ring")}
		{@render gridkey("E", "middle")}
		{@render gridkey("R", "index")}
		{@render gridkey("T", "index")}
		{@render gridkey("Y", "index")}
		{@render gridkey("U", "index")}
		{@render gridkey("I", "middle")}
		{@render gridkey("O", "ring")}
		{@render gridkey("P", "pinky")}
		{@render gridkey("[", "pinky")}
		{@render gridkey("]", "pinky")}
		{@render gridkey("\\", "pinky")}
		{@render gridkey("", "", false, true, false)}
		{@render gridkey("Del", "index")}
		{@render gridkey("End", "middle")}
		{@render gridkey("PgD", "ring")}
		{@render gridkey("", "", false, true, false)}
		{@render gridkey("7", "index")}
		{@render gridkey("8", "middle")}
		{@render gridkey("9", "ring")}
		<div class="key abs" class:active={isPressed("+")}>
			<p class:active={isPressed("+")}>+</p>
		</div>
		{@render gridkey("", "", false, false, false)}
	</div>
	<div class="row asdf">
		{@render gridkey("Caps", "pinky")}
		{@render gridkey("A", "pinky")}
		{@render gridkey("S", "ring")}
		{@render gridkey("D", "middle")}
		{@render gridkey("F", "index")}
		{@render gridkey("G", "index")}
		{@render gridkey("H", "index")}
		{@render gridkey("J", "index")}
		{@render gridkey("K", "middle")}
		{@render gridkey("L", "ring")}
		{@render gridkey(";", "pinky")}
		{@render gridkey("'", "pinky")}
		{@render gridkey("Enter", "pinky")}
		{@render gridkey("", "", false, true, false)}
		{@render gridkey("4", "index")}
		{@render gridkey("5", "middle")}
		{@render gridkey("6", "ring")}
		{@render gridkey("", "", false, false)}
	</div>
	<div class="row zxcv">
		{@render gridkey("Shift", "pinky")}
		{@render gridkey("Z", "pinky")}
		{@render gridkey("X", "ring")}
		{@render gridkey("C", "middle")}
		{@render gridkey("V", "index")}
		{@render gridkey("B", "index")}
		{@render gridkey("N", "index")}
		{@render gridkey("M", "index")}
		{@render gridkey(",", "middle")}
		{@render gridkey(".", "ring")}
		{@render gridkey("/", "pinky")}
		{@render gridkey("Shift", "pinky")}
		{@render gridkey("", "", false, false, false)}
		{@render gridkey("", "", false, true, true, false)}
		{@render gridkey("↑", "middle", true)}
		{@render gridkey("", "", false, false, true, false)}
		{@render gridkey("", "", false, true, false)}
		{@render gridkey("1", "index")}
		{@render gridkey("2", "middle")}
		{@render gridkey("3", "ring")}
		<div class="key abs" class:active={isPressed("E")}>
			<p class:active={isPressed("E")}>E</p>
		</div>
		{@render gridkey("", "", false, false, false)}
	</div>
	<div class="row bottom">
		{@render gridkey("Ctrl", "pinky")}
		{@render gridkey("⌘", "ring")}
		{@render gridkey("Alt", "ring")}
		{@render gridkey("Space", "thumb")}
		{@render gridkey("Alt", "ring")}
		{@render gridkey("⌘", "ring")}
		{@render gridkey("Fn", "ring")}
		{@render gridkey("Ctrl", "pinky")}
		{@render gridkey("", "")}
		{@render gridkey("←", "index")}
		{@render gridkey("↓", "middle")}
		{@render gridkey("→", "ring")}
		{@render gridkey("", "")}
		{@render gridkey("0", "index")}
		{@render gridkey(".", "ring")}
		{@render gridkey("", "", false, false)}
	</div>
</div>

<blockquote>this keyboard updates... try typing! •ᴗ•</blockquote>

<style>
	blockquote {
		padding-left: 1rem;
		margin-top: 2rem;
		border-left: 4px solid var(--surface1);
		color: var(--overlay1);
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
			color: var(--pink);
		}
		.key.ring {
			color: var(--sapphire);
		}
		.key.middle {
			color: var(--green);
		}
		.key.index {
			color: var(--peach);
		}
		.key.thumb {
			color: var(--red);
		}
		.key p.active {
			color: var(--rosewater);
			font-weight: 900;
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
