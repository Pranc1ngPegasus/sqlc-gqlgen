{ pkgs, lib, ... }:

{
  env = {
    PORT = "8080";
    DATABASE_HOST = "127.0.0.1";
    DATABASE_PORT = "5432";
    DATABASE_USER = "pranc1ngpegasus";
    DATABASE_PASSWORD = "password";
    DATABASE_NAME = "sqlc-gqlgen";
  };

  packages = with pkgs; [
    git
  ];

  scripts = { };

  enterShell = ''
    git --version
  '';

  languages = {
    nix.enable = true;
    go = {
      enable = true;
      package = pkgs.go_1_20;
    };
  };

  pre-commit = { };

  processes = { };

  services = {
    postgres = {
      enable = true;
      listen_addresses = "127.0.0.1";
      port = 5432;
      initialDatabases = [
        {
          name = "sqlc-gqlgen";
        }
      ];
      settings = {
        unix_socket_directories = lib.mkForce "";
      };
    };
  };
}
