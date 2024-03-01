CREATE MIGRATION m1houy3326i2owzeglczd4irmq5mtfb6bftq2xjumyijd7zhg26cha
    ONTO m1cicemi45ctkzkphzf4xonawjzqkqlhzbfz7av2tihrjyu4kzksea
{
  ALTER TYPE Users::Users {
      ALTER PROPERTY email {
          CREATE CONSTRAINT std::exclusive;
      };
      ALTER PROPERTY updated_at {
          SET readonly := true;
          CREATE REWRITE
              UPDATE 
              USING (std::datetime_of_statement());
      };
  };
};
