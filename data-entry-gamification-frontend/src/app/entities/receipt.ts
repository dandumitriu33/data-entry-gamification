export class Receipt {

    constructor(
      public id: number,
      public model_year: number,
      public make: string,
      public vin: string,
      public first_name: string,
      public last_name: string,
      public state: string
    ) {  }

  }