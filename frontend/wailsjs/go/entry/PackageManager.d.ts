// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {models} from '../models';

export function DisableApp(arg1:string):Promise<boolean>;

export function EnableApp(arg1:string):Promise<boolean>;

export function InstallExistingApp(arg1:string,arg2:Array<string>):Promise<boolean>;

export function ListApps(arg1:number):Promise<void>;

export function ListModuleApps(arg1:string,arg2:string):Promise<void>;

export function ListModules(arg1:string):Promise<Array<models.Module>>;

export function UninstallApp(arg1:string,arg2:Array<string>):Promise<boolean>;
