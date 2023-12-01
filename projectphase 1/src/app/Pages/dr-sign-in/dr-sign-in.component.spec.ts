import { ComponentFixture, TestBed } from '@angular/core/testing';

import { DrSignInComponent } from './dr-sign-in.component';

describe('DrSignInComponent', () => {
  let component: DrSignInComponent;
  let fixture: ComponentFixture<DrSignInComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [DrSignInComponent]
    });
    fixture = TestBed.createComponent(DrSignInComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
