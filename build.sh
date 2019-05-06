#!/bin/bash

FILE="/usr/include/gtk-3.0/gtk/gtkversion.h"
GTK_MAJOR=3
GTK_MINOR="$(grep '#define GTK_MINOR_VERSION' ${FILE} | grep -Po '[0-9]+')"
TAG="gtk_${GTK_MAJOR}_${GTK_MINOR}"

set -x

go build -tags "$TAG" $@