import UIKit

//MARK: - Class Definition
class ABrighterPlace{
    
    //MARK: - Properties
    var name: String
    var color: String
    var location: String
    
    //MARK: - Initialization
    init(name: String, color: String, location: String) {
        self.name = name
        self.color = color
        self.location = location
    }
    
    //MARK: - Instance Methods
    func printDetails() {
        print("The place \(name) has a \(color) color and is located at \(location)")
    }
    
    //MARK: - Class Methods
    static func loadPlaces() -> [ABrighterPlace] {
        // Sample Data
        let place1 = ABrighterPlace(name: "AmazingPark", color: "Blue", location: "5th Avenue")
        let place2 = ABrighterPlace(name: "CoolMountain", color: "Green", location: "6th Avenue")
        let place3 = ABrighterPlace(name: "LivelyBeach", color: "Yellow", location: "7th Avenue")
        
        return [place1, place2, place3]
    }
    
    static func savePlace(place: ABrighterPlace) {
        print("The place \(place.name) was added to the database!")
    }
    
    static func deletePlace(place: ABrighterPlace) {
        print("The place \(place.name) was deleted from the database!")
    }
    
    //MARK: - Nested Types
    enum PlaceData {
        case name, color, location
    }
}

//MARK: - Usage
let places = ABrighterPlace.loadPlaces()
for place in places {
    place.printDetails()
}

let newPlace = ABrighterPlace(name: "FantasticCity", color: "Purple", location: "8th Avenue")
ABrighterPlace.savePlace(place: newPlace)
ABrighterPlace.deletePlace(place: newPlace)

let placeDetails = ABrighterPlace.PlaceData.name
print(placeDetails)