{
  description = "a system-wide ascii keyboard visualizer in the terminal ";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
  };

  outputs = { self, nixpkgs }:
    let
      supportedSystems = [ "x86_64-linux" ];
      forAllSystems = nixpkgs.lib.genAttrs supportedSystems;
    in
    {
      packages = forAllSystems (system:
        let
          pkgs = nixpkgs.legacyPackages.${system};
        in
        {
          default = pkgs.buildGoModule {
            pname = "ditto";
            version = "1.0.5";

            src = ./.;

            vendorHash = "sha256-+DDBmGSsllHJ7D4/koKWq1MEVuUJJRebn3J8mxEQ8p8=";
          };
        });
    };
}
