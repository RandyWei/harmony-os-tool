export namespace models {
	
	export class App {
	    id: string;
	    name: string;
	    description: string;
	    related_ids: string[];
	    installed: boolean;
	    icon: string;
	    version: string;
	
	    static createFrom(source: any = {}) {
	        return new App(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.description = source["description"];
	        this.related_ids = source["related_ids"];
	        this.installed = source["installed"];
	        this.icon = source["icon"];
	        this.version = source["version"];
	    }
	}
	export class Device {
	    id: string;
	    product: string;
	    model: string;
	
	    static createFrom(source: any = {}) {
	        return new Device(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.product = source["product"];
	        this.model = source["model"];
	    }
	}
	export class EventData {
	    data: any;
	    type: number;
	
	    static createFrom(source: any = {}) {
	        return new EventData(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.data = source["data"];
	        this.type = source["type"];
	    }
	}
	export class Process {
	    pid: string;
	    virt: string;
	    res: string;
	    shr: string;
	    cpu: string;
	    mem: string;
	    time: string;
	    args: string;
	    desc: string;
	
	    static createFrom(source: any = {}) {
	        return new Process(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.pid = source["pid"];
	        this.virt = source["virt"];
	        this.res = source["res"];
	        this.shr = source["shr"];
	        this.cpu = source["cpu"];
	        this.mem = source["mem"];
	        this.time = source["time"];
	        this.args = source["args"];
	        this.desc = source["desc"];
	    }
	}
	export class TopInfo {
	    processes: Process[];
	    tasks: string;
	    mem: string;
	
	    static createFrom(source: any = {}) {
	        return new TopInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.processes = this.convertValues(source["processes"], Process);
	        this.tasks = source["tasks"];
	        this.mem = source["mem"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

