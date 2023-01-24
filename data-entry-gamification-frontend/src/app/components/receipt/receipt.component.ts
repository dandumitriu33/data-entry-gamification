import { Component, Input } from '@angular/core';
import { InterfaceReceipt } from 'src/app/interfaces/interface-receipt';
import { ReceiptService } from 'src/app/services/receipt.service';
import { ActivatedRoute } from '@angular/router';

@Component({
  selector: 'app-receipt',
  templateUrl: './receipt.component.html',
  styleUrls: ['./receipt.component.css']
})
export class ReceiptComponent {

  @Input() receipt?: InterfaceReceipt;

  ngOnInit(): void {
    this.getReceipt();
  }

  constructor(
    private receiptService: ReceiptService,
    private route: ActivatedRoute
  ) {}

  getReceipt(): void {
    const id = Number(this.route.snapshot.paramMap.get('id'));
    this.receiptService.getReceipt(id)
      .subscribe(receipt => this.receipt = receipt);
  }

}
