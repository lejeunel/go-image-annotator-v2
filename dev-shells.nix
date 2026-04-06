{
  perSystem =
    { pkgs, ... }:
    {

      devShells.default = pkgs.mkShell {
        buildInputs = with pkgs; [
          gopls
          gotest
          oapi-codegen
          tailwindcss_4
        ];
      };
    };

}
