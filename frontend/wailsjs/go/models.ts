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
	export class Category {
	    ID: number;
	    Name: string;
	    Type: string;
	
	    static createFrom(source: any = {}) {
	        return new Category(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ID = source["ID"];
	        this.Name = source["Name"];
	        this.Type = source["Type"];
	    }
	}
	export class CreateTransactionRequest {
	    Date: string;
	    Amount: number;
	    Comment: string;
	    CategoryID: number;
	
	    static createFrom(source: any = {}) {
	        return new CreateTransactionRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Date = source["Date"];
	        this.Amount = source["Amount"];
	        this.Comment = source["Comment"];
	        this.CategoryID = source["CategoryID"];
	    }
	}
	export class Transaction {
	    ID: number;
	    Date: string;
	    Amount: string;
	    Comment: string;
	    IsIncome: boolean;
	    Category: string;
	
	    static createFrom(source: any = {}) {
	        return new Transaction(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ID = source["ID"];
	        this.Date = source["Date"];
	        this.Amount = source["Amount"];
	        this.Comment = source["Comment"];
	        this.IsIncome = source["IsIncome"];
	        this.Category = source["Category"];
	    }
	}

}

