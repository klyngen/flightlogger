package storage

import (
	"github.com/jinzhu/gorm"
	"github.com/klyngen/flightlogger/common"
)

// ######################## MAP FROM BUSINESS / PRESENTATION INTO DATA #########################

func mapUser(user common.User) DbUser {
	return DbUser{
		Model: gorm.Model{
			ID: user.ID,
		},
		Username:  user.Username,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Clubs:     mapClubs(user.Clubs),
		Groups:    mapUserGroups(user.Groups),
		Credentials: DbCredentials{
			PasswordHash: user.PasswordHash,
			PasswordSalt: user.PasswordSalt,
		},
		Scopes: mapScopes(user.Scopes),
		Wings:  mapWings(user.Wings),
	}
}

func mapScopes(scopes []common.UserScope) []DbUserScope {
	newScopes := make([]DbUserScope, len(scopes))
	for i, s := range scopes {
		newScopes[i] = mapScope(s)
	}
	return newScopes
}

func mapScope(scope common.UserScope) DbUserScope {
	return DbUserScope{
		Key:  scope.Key,
		Name: scope.Name,
		Model: gorm.Model{
			ID: scope.ID,
		},
	}
}

func mapClub(club common.Club) DbClub {
	return DbClub{
		Description: club.Description,
		Name:        club.Name,
		Model: gorm.Model{
			ID: club.ID,
		},
		Location: mapLocation(club.Location),
	}
}

func mapClubs(clubs []common.Club) []DbClub {
	newClubs := make([]DbClub, len(clubs))

	for i, c := range clubs {
		newClubs[i] = mapClub(c)
	}

	return newClubs
}

func mapUserGroup(group common.UserGroup) DbUserGroup {
	return DbUserGroup{
		Model: gorm.Model{
			ID: group.ID,
		},
		Name:   group.Name,
		Key:    group.Key,
		Scopes: mapScopes(group.Scopes),
	}
}

func mapUserGroups(groups []common.UserGroup) []DbUserGroup {
	newGroups := make([]DbUserGroup, len(groups))

	for i, g := range groups {
		newGroups[i] = mapUserGroup(g)
	}

	return newGroups
}

func mapWing(wing common.Wing) DbWing {
	return DbWing{
		Model: gorm.Model{
			ID: wing.ID,
		},
		Name: wing.Name,
		Details: DbWingScoreDetails{
			DhvScore: wing.Details.DhvScore,
			EnaScore: wing.Details.EnaScore,
			Model: gorm.Model{
				ID: wing.Details.ID,
			},
			Description: wing.Details.Description,
		},
		Images: mapFileReferences(wing.Images),
	}
}

func mapWings(wing []common.Wing) []DbWing {
	newWings := make([]DbWing, len(wing))

	for i, w := range wing {
		newWings[i] = mapWing(w)
	}

	return newWings
}

func mapFileReference(file common.FileReference) DbFileReference {
	return DbFileReference{
		Model: gorm.Model{
			ID: file.ID,
		},
		MimeType:     file.MimeType,
		FileName:     file.FileName,
		FileLocation: file.FileLocation,
	}
}

func mapFileReferences(file []common.FileReference) []DbFileReference {
	newFiles := make([]DbFileReference, len(file))

	for i, f := range file {
		newFiles[i] = mapFileReference(f)
	}

	return newFiles
}

func mapLocation(loc common.Location) DbLocation {
	return DbLocation{
		Model: gorm.Model{
			ID: loc.ID,
		},
		Name:        loc.Name,
		Description: loc.Description,
		Elevation:   loc.Elevation,
		Coordinates: DbCoordinates{
			Longitude: loc.Longitude,
			Lattitude: loc.Lattitude,
		},
	}
}

func mapStartsite(site common.StartSite) DbStartSite {
	return DbStartSite{
		Model: gorm.Model{
			ID: site.ID,
		},
		Description: site.Description,
		Difficulty:  site.Difficulty,
		Location:    mapLocation(site.Location),
		Waypoints:   mapWaypoints(site.Waypoints),
	}
}

func mapWayPoint(waypoint common.Waypoint) DbWaypoint {
	return DbWaypoint{
		Difficulty: waypoint.Difficulty,
		Location:   mapLocation(waypoint.Location),
		Model: gorm.Model{
			ID: waypoint.ID,
		},
	}
}

func mapWaypoints(waypoints []common.Waypoint) []DbWaypoint {
	newPoints := make([]DbWaypoint, len(waypoints))

	for i, w := range waypoints {
		newPoints[i] = mapWayPoint(w)
	}

	return newPoints
}

func mapFlight(flight common.Flight) DbFlight {
	return DbFlight{
		Model: gorm.Model{
			ID: flight.ID,
		},
		Startsite: mapStartsite(flight.Startsite),
		User:      mapUser(flight.User),
		Waypoint:  mapWayPoint(flight.Waypoint),
		Duration:  flight.Duration,
		Notes:     flight.Notes,
		Distance:  flight.Distance,
		MaxHight:  flight.MaxHight,
		HangTime:  flight.HangTime,
		Wing:      mapWing(flight.Wing),
		Photos:    mapFileReferences(flight.Photos),
		FlightLog: mapFileReference(flight.FlightLog),
		FlightType: DbFlightType{
			Model: gorm.Model{
				ID: flight.FlightType.Value,
			},
			Name: flight.FlightType.Name,
		},
		TakeOffType: DbTakeoffType{
			Model: gorm.Model{
				ID: flight.TakeOffType.Value,
			},
			Name: flight.TakeOffType.Name,
		},
		Incidents: mapIncidents(flight.Incidents),
	}
}

func mapIncident(incident common.Incident) DbIncident {
	return DbIncident{
		Model: gorm.Model{
			ID: incident.ID,
		},
		Level:             incident.Level,
		Description:       incident.Description,
		Public:            incident.Public,
		NotifiedAmbulance: incident.NotifiedAmbulance,
		NotifiedPolice:    incident.NotifiedPolice,
		LatestFlight: DbFlight{
			Model: gorm.Model{
				ID: incident.LatestFlightID,
			},
		},
		Weatherconfitions: incident.Weatherconfitions,
	}
}

func mapIncidents(incidents []common.Incident) []DbIncident {
	newInc := make([]DbIncident, len(incidents))

	for i, in := range incidents {
		newInc[i] = mapIncident(in)
	}

	return newInc
}

// ######################## MAP FROM DATA LAYER TO PRESENTATION / BUSINESS #########################

func demapUser(user DbUser) common.User {
	return common.User{
		ID:        user.Model.ID,
		Username:  user.Username,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Scopes:    demapScopes(user.Scopes),
		Clubs:     demapClubs(user.Clubs),
		Groups:    demapGroups(user.Groups),
		Wings:     demapWings(user.Wings),
	}
}

func demapUsers(users []DbUser) []common.User {
	newUsers := make([]common.User, len(users))

	for i, u := range users {
		newUsers[i] = demapUser(u)
	}

	return newUsers
}

func demapScope(scope DbUserScope) common.UserScope {
	return common.UserScope{
		Key:  scope.Key,
		Name: scope.Name,
		ID:   scope.Model.ID,
	}
}

func demapScopes(scopes []DbUserScope) []common.UserScope {
	newScopes := make([]common.UserScope, len(scopes))

	for i, r := range scopes {
		newScopes[i] = demapScope(r)
	}

	return newScopes
}

func demapClub(club DbClub) common.Club {
	return common.Club{
		Description: club.Description,
		ID:          club.Model.ID,
		Name:        club.Name,
		Location:    demapLocation(club.Location),
	}
}

func demapClubs(club []DbClub) []common.Club {
	newClubs := make([]common.Club, len(club))

	for i, c := range club {
		newClubs[i] = demapClub(c)
	}

	return newClubs
}

func demapLocation(club DbLocation) common.Location {
	return common.Location{
		Lattitude:   club.Coordinates.Lattitude,
		Longitude:   club.Coordinates.Longitude,
		Name:        club.Name,
		Elevation:   club.Elevation,
		Description: club.Description,
		ID:          club.Model.ID,
	}
}

func demapGroup(group DbUserGroup) common.UserGroup {
	return common.UserGroup{
		ID:     group.Model.ID,
		Key:    group.Key,
		Name:   group.Name,
		Scopes: demapScopes(group.Scopes),
	}
}

func demapGroups(groups []DbUserGroup) []common.UserGroup {
	newGroups := make([]common.UserGroup, len(groups))

	for i, g := range groups {
		newGroups[i] = demapGroup(g)
	}

	return newGroups
}

func demapWing(wing DbWing) common.Wing {
	return common.Wing{
		ID:   wing.Model.ID,
		Name: wing.Name,
		Details: common.WingDetails{
			DhvScore:    wing.Details.DhvScore,
			EnaScore:    wing.Details.EnaScore,
			Description: wing.Details.Description,
			ID:          wing.Details.Model.ID,
		},
		Images: demapFileReferences(wing.Images),
	}
}

func demapWings(wings []DbWing) []common.Wing {
	newWings := make([]common.Wing, len(wings))

	for i, w := range wings {
		newWings[i] = demapWing(w)
	}

	return newWings
}

func demapFileReference(file DbFileReference) common.FileReference {
	return common.FileReference{
		ID:           file.Model.ID,
		MimeType:     file.MimeType,
		FileName:     file.FileName,
		FileLocation: file.FileLocation,
	}
}

func demapFileReferences(file []DbFileReference) []common.FileReference {
	newFiles := make([]common.FileReference, len(file))

	for i, f := range file {
		newFiles[i] = demapFileReference(f)
	}

	return newFiles
}

func demapStartSite(site DbStartSite) common.StartSite {
	return common.StartSite{
		Difficulty:  site.Difficulty,
		Description: site.Description,
		ID:          site.ID,
		Location:    demapLocation(site.Location),
		Waypoints:   demapWaypoints(site.Waypoints),
	}
}

func demapWaypoint(point DbWaypoint) common.Waypoint {
	return common.Waypoint{
		Difficulty: point.Difficulty,
		Location:   demapLocation(point.Location),
		ID:         point.ID,
	}
}

func demapWaypoints(points []DbWaypoint) []common.Waypoint {
	newPoints := make([]common.Waypoint, len(points))

	for i, p := range points {
		newPoints[i] = demapWaypoint(p)
	}

	return newPoints
}

func demapFlight(flight DbFlight) common.Flight {
	return common.Flight{
		ID:        flight.ID,
		Startsite: demapStartSite(flight.Startsite),
		User:      demapUser(flight.User),
		Waypoint:  demapWaypoint(flight.Waypoint),
		Duration:  flight.Duration,
		MaxHight:  flight.MaxHight,
		HangTime:  flight.HangTime,
		Wing:      demapWing(flight.Wing),
		Photos:    demapFileReferences(flight.Photos),
		FlightLog: demapFileReference(flight.FlightLog),
		FlightType: common.FlightType{
			Value: flight.FlightType.ID,
			Name:  flight.FlightType.Name,
		},
		TakeOffType: common.TakeoffType{
			Value: flight.TakeOffType.ID,
			Name:  flight.TakeOffType.Name,
		},
		Incidents: demapIncidents(flight.Incidents),
	}
}

func demapIncident(incident DbIncident) common.Incident {
	return common.Incident{
		ID:                incident.ID,
		Level:             incident.Level,
		Description:       incident.Description,
		Public:            incident.Public,
		NotifiedAmbulance: incident.NotifiedAmbulance,
		NotifiedPolice:    incident.NotifiedPolice,
		LatestFlightID:    incident.LatestFlight.ID,
		Weatherconfitions: incident.Weatherconfitions,
	}
}

func demapIncidents(incidents []DbIncident) []common.Incident {
	newIncidents := make([]common.Incident, len(incidents))

	for i, in := range incidents {
		newIncidents[i] = demapIncident(in)
	}

	return newIncidents
}
