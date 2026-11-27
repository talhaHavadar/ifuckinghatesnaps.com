{
  description = "ifuckinghatesnaps.com development environment";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs =
    { nixpkgs, flake-utils, ... }:
    flake-utils.lib.eachDefaultSystem (
      system:
      let
        pkgs = nixpkgs.legacyPackages.${system};
      in
      {
        devShells.default = pkgs.mkShell {
          packages = with pkgs; [
            go
            gopls
            golangci-lint
            govulncheck

            just
            prek
            shellcheck
            hadolint
            actionlint
            curl
          ];

          # Use the Go from this shell, never a toolchain downloaded because of go.mod.
          GOTOOLCHAIN = "local";
        };
      }
    );
}
