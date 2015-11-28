#include <regex>
#include <sstream>
#include "b64/decode.h"
#include "Module/Base64.h"

namespace S7_Tfit_Module {

std::string Base64::get_display_name() {
    return "Base64";
}

std::string Base64::match(std::string line) {
    if (!regex_match(line, std::regex("^(?:[a-zA-Z0-9+/]{4})*(?:[a-zA-Z0-9+/]{2}==|[a-zA-Z0-9+/]{3}=)?$"))) {
        return "";
    }
    std::stringstream input {line};
    std::stringstream output;
    base64::decoder dec;
    dec.decode(input, output);
    return output.str();
}

}
