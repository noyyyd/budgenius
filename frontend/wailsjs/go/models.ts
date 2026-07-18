export namespace main {
	
	export class Budget {
	    id: number;
	    name: string;
	    start: string;
	    end: string;
	
	    static createFrom(source: any = {}) {
	        return new Budget(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.start = source["start"];
	        this.end = source["end"];
	    }
	}

}

