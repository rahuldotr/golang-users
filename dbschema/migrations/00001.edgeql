CREATE MIGRATION m1cicemi45ctkzkphzf4xonawjzqkqlhzbfz7av2tihrjyu4kzksea
    ONTO initial
{
  CREATE MODULE Users IF NOT EXISTS;
  CREATE TYPE Users::Users {
      CREATE PROPERTY bio: std::str;
      CREATE PROPERTY company: std::str;
      CREATE REQUIRED PROPERTY created_at: std::datetime {
          SET default := (std::datetime_of_statement());
          SET readonly := true;
      };
      CREATE REQUIRED PROPERTY email: std::str;
      CREATE PROPERTY location: std::str;
      CREATE REQUIRED PROPERTY name: std::str;
      CREATE PROPERTY twitter_username: std::str;
      CREATE PROPERTY updated_at: std::datetime;
  };
};
