CREATE MIGRATION m1jxfqcxop5zr42loma6boev3bwux5ztaqiixprnvfv6auddy5b6ea
    ONTO m1houy3326i2owzeglczd4irmq5mtfb6bftq2xjumyijd7zhg26cha
{
  ALTER TYPE Users::Users {
      ALTER PROPERTY updated_at {
          RESET readonly;
      };
  };
};
