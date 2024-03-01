module Users {
    type Users {
        required name: str;
        company: str;
        location: str;
        required email: str {
            constraint exclusive;
        };
        bio: str;
        twitter_username: str;
        required created_at: datetime {
            readonly := true;
            default := datetime_of_statement();
        };
        updated_at: datetime {
            rewrite update using (datetime_of_statement());
        };
    }
}