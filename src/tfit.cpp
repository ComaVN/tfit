/*
 * main.cpp
 * Main function for TFIT executable
 */

#include <getopt.h>
#include <iostream>
#include <string>
#include "cmdline.h"
#include "Module/Base64.h"

int main(int argc, char *argv[]) {
    gengetopt_args_info args_info;

    if (cmdline_parser(argc, argv, &args_info) != 0) {
        exit(1);
    }

    S7_Tfit_Module::IModule* modules[1] = {
        new S7_Tfit_Module::Base64(),
    };

    std::string line;
    std::cin >> line;
    for (int i=0; i<1; ++i) {
        std::string result = modules[i]->match(line);
        if (result.length() > 0) {
            std::cout << modules[i]->get_display_name() << " matched:\n" << result;
        }
    }
    std::cout << "\n";

    for (int i=0; i<1; ++i) {
        delete modules[i];
    }
}
