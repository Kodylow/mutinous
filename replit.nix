{ pkgs }: {
    deps = [
        pkgs.jq.bin
        pkgs.clightning
        pkgs.just
        pkgs.go
        pkgs.gopls
    ];
}