#include <SPI.h>
#include "Adafruit_MAX31855.h"

// Default connection is using software SPI, but comment and uncomment one of
// the two examples below to switch between software SPI and hardware SPI:

// Example creating a thermocouple instance with software SPI on any three
// digital IO pins.
#define MAXDO   2
#define MAXCS   3
#define MAXCLK  4

#define measurement_time 5000

#define MAXDO2   11
#define MAXCS2   12
#define MAXCLK2  13

// initialize the Thermocouples
Adafruit_MAX31855 thermocouple(MAXCLK, MAXCS, MAXDO);
Adafruit_MAX31855 thermocouple2(MAXCLK2, MAXCS2, MAXDO2);

// Example creating a thermocouple instance with hardware SPI (Uno/Mega only)
// on a given CS pin.
//#define MAXCS   10
//Adafruit_MAX31855 thermocouple(MAXCS);

void setup() {  
  Serial.begin(9600);
  
  Serial.println("Time\tAmbient\tC1\tC2");
  // wait for MAX chip to stabilize
  delay(500);
}

unsigned long loop_count = 0;

void loop() {
  loop_count = (loop_count + 1); 
  while(loop_count*measurement_time - 270 >millis()){
     delay(1);
   }
   // basic readout test, just print the current temp
   double a1 = (thermocouple.readInternal());
   
   double a2 = (thermocouple2.readInternal());
   
   double c1 = thermocouple.readCelsius();
   if (isnan(c1)) {
     Serial.println("Something wrong with thermocouple!");
   }
   
   double c2 = thermocouple2.readCelsius();
   if (isnan(c2)) {
     Serial.println("Something wrong with thermocouple2!");
   }
   
   //display / print
   Serial.print(millis()/1000.0);
   Serial.print("\t");
   Serial.print((a1+a2)/2);
   Serial.print("\t");
   Serial.print(c1);
   Serial.print("\t");
   Serial.print(c2);
   Serial.println();
   
   
   //Serial.print("F = ");
   //Serial.println(thermocouple.readFarenheit());
 
}
