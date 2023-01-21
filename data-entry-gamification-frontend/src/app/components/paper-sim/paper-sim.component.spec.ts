import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PaperSimComponent } from './paper-sim.component';

describe('PaperSimComponent', () => {
  let component: PaperSimComponent;
  let fixture: ComponentFixture<PaperSimComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ PaperSimComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(PaperSimComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
