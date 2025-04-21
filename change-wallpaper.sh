dbus-send --session --dest=org.kde.plasmashell --type=method_call /PlasmaShell org.kde.PlasmaShell.evaluateScript 'string:
var Desktops = desktops();                                                                                                                       
d = Desktops[0];
d.wallpaperPlugin = "org.kde.image";
d.currentConfigGroup = Array("Wallpaper",
                         "org.kde.image",
                         "General");
d.writeConfig("Image", "file:///home/tainticulus/Documents/wallpapers/image-horizontal.jpg");


d = Desktops[1];
d.wallpaperPlugin = "org.kde.image";
d.currentConfigGroup = Array("Wallpaper",
                         "org.kde.image",
                         "General");
d.writeConfig("Image", "file:///home/tainticulus/Documents/wallpapers/image-vertical.jpg");
'
