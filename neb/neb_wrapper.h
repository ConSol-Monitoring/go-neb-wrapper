#include "dependencies.h"

/* This file contains the glue for calling go functions from
 * within the naemon core.
 */

NEB_API_VERSION( CURRENT_NEB_API_VERSION );

nebmodule *neb_handle;

extern int GoNebModuleInit(int flags, char *args);
extern int GoNebModuleDeinit(int flags, int reason);
extern int Generic_Callback(int type, void* data);

int nebmodule_init( int flags, char *args, nebmodule *handle ) {
    neb_handle = handle;

    return(GoNebModuleInit(flags, args));
}

int nebmodule_deinit( int flags, int reason ) {
    return(GoNebModuleDeinit(flags, reason));
}

int generic_callback(int type,  void* data) {
   return(Generic_Callback(type, data));
}

