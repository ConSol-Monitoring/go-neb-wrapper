#define CORE_NAGIOS3    1
#define CORE_NAGIOS4    2
#define CORE_NAEMON     3

int CORE_TYPE;

#if defined(NAGIOS3)
#include "nagios3/nagios.h"
#include "nagios3/neberrors.h"
#include "nagios3/nebstructs.h"
#include "nagios3/nebcallbacks.h"
#include "nagios3/broker.h"
#include "nagios3/macros.h"
#include "nagios3/nebcallbacks.h"
#elif defined(NAGIOS4)
#include "nagios4/nagios.h"
#include "nagios4/neberrors.h"
#include "nagios4/nebstructs.h"
#include "nagios4/nebcallbacks.h"
#include "nagios4/broker.h"
#include "nagios4/macros.h"
#include "nagios4/nebcallbacks.h"
#elif defined(NAEMON)
#include "naemon/naemon.h"
#else
#error "must specify one of NAGIOS3, NAGIOS4 or NAEMON"
#endif