CXX = clang++
CXXFLAGS = -std=c++11 -stdlib=libc++
GENGETOPT = gengetopt
SOURCEDIR = src
BUILDDIR = build
EXECUTABLE = tfit
CPPFILES = $(SOURCEDIR)/$(EXECUTABLE).cpp $(SOURCEDIR)/cmdline.cpp $(wildcard $(SOURCEDIR)/**/*.cpp)

all: $(BUILDDIR)/$(EXECUTABLE)

$(BUILDDIR):
	mkdir -p $(BUILDDIR)

$(BUILDDIR)/$(EXECUTABLE): $(BUILDDIR) $(CPPFILES)
	$(CXX) $(CXXFLAGS) ${CPPFILES} -o $(BUILDDIR)/$(EXECUTABLE) -I$(SOURCEDIR)

$(SOURCEDIR)/cmdline.cpp: $(SOURCEDIR)/$(EXECUTABLE).ggo
	$(GENGETOPT) --output-dir $(SOURCEDIR) --c-extension cpp < $(SOURCEDIR)/$(EXECUTABLE).ggo

clean:
	rm -rf $(BUILDDIR) $(SOURCEDIR)/cmdline.cpp $(SOURCEDIR)/cmdline.h

.PHONY: all gengetopt clean
