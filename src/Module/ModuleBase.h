#ifndef S7_TFIT_MODULE_MODULEBASE_H_INCLUDED
#define S7_TFIT_MODULE_MODULEBASE_H_INCLUDED

#include <string>

namespace S7_Tfit_Module {

class ModuleBase {

    public:
        virtual std::string match(std::string line)=0;

};

}

#endif
