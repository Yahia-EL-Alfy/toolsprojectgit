import { ComponentFixture, TestBed } from '@angular/core/testing';

import { DrSlotsComponent } from './dr-slots.component';

describe('DrSlotsComponent', () => {
  let component: DrSlotsComponent;
  let fixture: ComponentFixture<DrSlotsComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [DrSlotsComponent]
    });
    fixture = TestBed.createComponent(DrSlotsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
