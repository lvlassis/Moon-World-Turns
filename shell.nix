{ pkgs ? import <nixpkgs> {} }:

pkgs.mkShell {
  buildInputs = with pkgs; [
    # Go
    go

    # Graphics libraries
    xorg.libX11
    xorg.libXcursor
    xorg.libXrandr
    xorg.libXinerama
    xorg.libXi
    xorg.libXxf86vm

    # OpenGL
    libGL
    libGLU

    # Audio libraries
    openal
    libvorbis
    libogg

    # GLFW dependencies
    glfw

    # Build tools
    pkg-config
    gcc
  ];

  shellHook = ''
    export CGO_CFLAGS="-I${pkgs.openal}/include/AL -I${pkgs.libvorbis.dev}/include/vorbis"
    export CGO_LDFLAGS="-L${pkgs.openal}/lib -L${pkgs.libvorbis}/lib -L${pkgs.libogg}/lib"
    echo "Ambiente de desenvolvimento g3n configurado!"
    echo "Caminho libvorbis: ${pkgs.libvorbis.dev}/include"
    echo "Agora você pode executar: go run ."
  '';

  # Set CGO flags to help the compiler find libraries
  CGO_ENABLED = "1";
}
