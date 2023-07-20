{ pkgs }: {
    deps = [
        pkgs.clightning
        pkgs.just
        pkgs.go
        pkgs.gopls
    ];
}