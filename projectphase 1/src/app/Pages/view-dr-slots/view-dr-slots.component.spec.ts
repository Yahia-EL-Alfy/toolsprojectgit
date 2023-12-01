import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ViewDrSlotsComponent } from './view-dr-slots.component';

describe('ViewDrSlotsComponent', () => {
  let component: ViewDrSlotsComponent;
  let fixture: ComponentFixture<ViewDrSlotsComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [ViewDrSlotsComponent]
    });
    fixture = TestBed.createComponent(ViewDrSlotsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
