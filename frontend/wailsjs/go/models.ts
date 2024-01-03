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

}

