/*
 * main.cpp
 * Main function for TFIT executable
 */

#include <getopt.h>
#include <iostream>
#include <string>
#include <vector>
#include "cmdline.h"
#include "Module/Base64.h"

int main(int argc, char *argv[]) {
    gengetopt_args_info args_info;

    if (cmdline_parser(argc, argv, &args_info) != 0) {
        exit(1);
    }

    std::vector<S7_Tfit_Module::IModule*> modules;
    modules.push_back(new S7_Tfit_Module::Base64());

    std::string line;
    std::cin >> line;
    for (auto& module : modules) {
        std::string result = module->match(line);
        if (result.length() > 0) {
            std::cout << module->get_display_name() << " matched:\n" << result;
        }
    }
    std::cout << "\n";

    for (auto& module : modules) {
        delete module;
    }
}
