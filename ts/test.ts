type CountryCode = "FR" | "ES" | "BE";
type RegionCode = `${CountryCode}-${string}`;

const countryCode: CountryCode = "FR";

const regionCode: RegionCode = "FR-SIS";
