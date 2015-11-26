#ifndef S7_TFIT_MODULE_BASE64_H_INCLUDED
#define S7_TFIT_MODULE_BASE64_H_INCLUDED

#include "Module/ModuleBase.h"

namespace S7_Tfit_Module {

class Base64: public ModuleBase {

    public:
        std::string match(std::string line) override;

};

}

#endif
