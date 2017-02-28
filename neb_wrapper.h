#include "naemon/naemon.h"

/* This file contains the glue for calling go functions from
 * within the naemon core.
 */

NEB_API_VERSION( CURRENT_NEB_API_VERSION );

nebmodule *neb_handle;

extern int Neb_Module_Init(int flags, char *args);
extern int Neb_Module_Deinit(int flags, int reason);
extern int Process_Data_Callback(int type, void* data);

int nebmodule_init( int flags, char *args, nebmodule *handle ) {
    neb_handle = handle;
    return(Neb_Module_Init(flags, args));
}

int nebmodule_deinit( int flags, int reason ) {
    return(Neb_Module_Deinit(flags, reason));
}

int process_data_callback(int type,  void* data) {
   return(Process_Data_Callback(type, data));
}

