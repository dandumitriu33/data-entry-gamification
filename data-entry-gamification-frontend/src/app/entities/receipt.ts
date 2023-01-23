export class Receipt {

    constructor(
      public id: number,
      public modelYear: number,
      public make: string,
      public vin: string,
      public firstName: string,
      public lastName: string,
      public state: string
    ) {  }

  }