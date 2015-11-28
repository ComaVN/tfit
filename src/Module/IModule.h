#ifndef S7_TFIT_MODULE_IMODULE_H_INCLUDED
#define S7_TFIT_MODULE_IMODULE_H_INCLUDED

#include <string>

namespace S7_Tfit_Module {

class IModule {

    public:
        virtual std::string get_display_name() = 0;
        virtual std::string match(std::string line) = 0;
        virtual ~IModule() {};
};

}

#endif
