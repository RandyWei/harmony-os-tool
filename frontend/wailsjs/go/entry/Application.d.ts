// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {models} from '../models';

export function GetAppDir():Promise<string>;

export function Greet(arg1:string):Promise<string>;

export function InstallAdb():Promise<boolean>;

export function InstallExistingApp(arg1:string):Promise<boolean>;

export function ListApps():Promise<Array<models.App>>;

export function UninstallApp(arg1:string):Promise<boolean>;

export function WaitForDevice():Promise<Array<models.Device>>;
