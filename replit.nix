{ pkgs }: {
    deps = [
        pkgs.pstree
        pkgs.jq.bin
        pkgs.clightning
        pkgs.just
        pkgs.go
        pkgs.gopls
    ];
}