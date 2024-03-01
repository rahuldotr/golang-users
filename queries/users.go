package queries

func GetInsertUserQuery() string {
	query := `INSERT Users::Users {
        name := <str>$name,
        email := <str>$email,
        company := <optional str>$company,
        bio := <optional str>$bio,
        location := <optional str>$location,
        twitter_username := <optional str>$twitter_username
    };`
	return query
}

func GetAllUsersQuery() string {
	query := `SELECT Users::Users { id, name, email, company, bio, location, twitter_username };`
	return query
}

func GetIndividualUserQuery() string {
	query := `SELECT Users::Users { id, name, email, company, bio, location, twitter_username } FILTER .id = <uuid>$id;`
	return query
}

func UpdateUserQuery() string {
	query := `UPDATE Users::Users FILTER .id = <uuid>$id SET {
		name := <str>$name, 
		email := <str>$email, 
		company := <optional str>$company,
        bio := <optional str>$bio,
        location := <optional str>$location,
        twitter_username := <optional str>$twitter_username
	};`
	return query
}

func DeleteUserQuery() string {
	query := `DELETE Users::Users filter .id = <uuid>$id ;`
	return query
}

func GetUserWithEmailCountQuery() string {
	query := `SELECT count((SELECT Users::Users FILTER .email = <str>$email));`
	return query
}

func CheckEmailAssignedToOtherUser() string {
	query := `SELECT count((SELECT Users::Users FILTER .email = <str>$email and .id != <uuid>$id));`
	return query
}
