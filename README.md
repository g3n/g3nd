# G3ND - G3N Game Engine Demo Program

G3ND is a demo/test program for the [G3N](https://github.com/g3n/engine) Go 3D Game Engine.
It contains demos of the main features of the engine and also some basic tests.
It can also be used to learn how to use the game engine by examining the source code of the demo programs.
It is very easy to create a new demo or test as the main program takes care
of a lot of necessary initializations and house keeping.

<p align="center">
  <img style="float: right;" src="data/images/cube.png" alt="G3ND Screenshot"/>
</p>

# Dependencies for installation

G3ND imports the [G3N](https://github.com/g3n/engine) and so has the same dependencies as the engine itself.
It needs an OpenGL driver installed in your system and
on Unix like systems depends on some C libraries that must be installed.
In all cases it is necessary to have a C compiler installed.

* For Ubuntu/Debian-like Linux distributions, install `libgl1-mesa-dev` and `xorg-dev` packages.
* For CentOS/Fedora-like Linux distributions, install `libX11-devel libXcursor-devel libXrandr-devel libXinerama-devel mesa-libGL-devel libXi-devel` packages.
* Currently it was not tested on OS X. We encourage some feedback.
* For Windows it is recommended to have MingGW vx.x.x installed.

G3ND checks if audio libraries are installed in the system at runtime
and if found enables the execution of audio demos.
The following libraries are necessary for the audio demos:

* For Ubuntu/Debian-like Linux distributions, install `libopenal1` and `libvorbisfile3`
* For CentOS/Fedora-like Linux distributions, install `libopenal1` and `libvorbisfile3`
* Currently it was not tested on OS X. We encourage some feedback.
* For Windows install `OpenAL32.dll` and ???

G3ND was only tested with Go1.7.4+.

# Installation

`go get -u github.com/g3n/g3nd`

Note: G3ND comes with a data directory with media files: images, textures, models and audio files.
Currently this directory has aproximately 50MB.

# Running

When G3ND is run without any command line parameters it shows the tree of
categorized available demos at the left of its window and an empty center area
to show the demo scene.
Click on a category in the tree to expand it and then select a demo to show.

At the upper right corner is located the `Control` folder, which when clicked
shows some controls which can change the parameters of the current demo.

To exit the program press ESC or close the window.

You can start G3ND to show a specific demo specifying the demo name (category plus underscore plus name) in the command
line such as:

`>g3nd geometry_box`

To check the maximum frames per second rate (FPS) of your system for any demo,
run G3ND with the option `-interval 0`.
Note that at least one core of your system CPU will run at 100% utilization.
The FPS will be lower when the screen is maximized or full.

# Creating a new demo/test

To create a new demo or test create a file in G3ND's main directory
with a name like `category_name.go`. For example: 'tests_hello.go'

# Contributing

If you spot a bug or create a new interesting demo you are encouraged to
send pull requests.


