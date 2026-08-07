export namespace traversal {
	
	export class Option {
	    Id: string;
	    Text: string;
	    PreviewText: string;
	    Requirements: Record<string, number>;
	
	    static createFrom(source: any = {}) {
	        return new Option(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Id = source["Id"];
	        this.Text = source["Text"];
	        this.PreviewText = source["PreviewText"];
	        this.Requirements = source["Requirements"];
	    }
	}

}

