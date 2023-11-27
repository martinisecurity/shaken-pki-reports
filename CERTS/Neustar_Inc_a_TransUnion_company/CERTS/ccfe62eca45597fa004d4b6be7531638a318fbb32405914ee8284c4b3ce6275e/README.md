# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Neustar UAT Enterprise Certified Caller Intermediate CA1

Tested At: 27 Nov 23 23:28 UTC\
Initial Validity Period: 3653 day(s)\
Remaining Validity Period: 2983 day(s)\
Subject: CN=Neustar UAT Enterprise Certified Caller Intermediate CA1, OU=www.ccid-uat.neustar, O=Neustar Inc a TransUnion company, C=US\
Issuer: CN=Neustar UAT Enterprise Certified Caller Root CA, OU=www.ccid-uat.neustar, O=Neustar Inc a TransUnion company, C=US

[View certificate details](https://understandingwebpki.com/?cert=MIIDETCCAregAwIBAgIUMeczGmSlM1e%2BQ6bA0Od9W23%2BXKowCgYIKoZIzj0EAwIwgZExCzAJBgNVBAYTAlVTMSkwJwYDVQQKDCBOZXVzdGFyIEluYyBhIFRyYW5zVW5pb24gY29tcGFueTEdMBsGA1UECwwUd3d3LmNjaWQtdWF0Lm5ldXN0YXIxODA2BgNVBAMML05ldXN0YXIgVUFUIEVudGVycHJpc2UgQ2VydGlmaWVkIENhbGxlciBSb290IENBMB4XDTIyMDEyNjE1NTQ1MVoXDTMyMDEyNzE1NTQ1MVowgZoxCzAJBgNVBAYTAlVTMSkwJwYDVQQKDCBOZXVzdGFyIEluYyBhIFRyYW5zVW5pb24gY29tcGFueTEdMBsGA1UECwwUd3d3LmNjaWQtdWF0Lm5ldXN0YXIxQTA%2FBgNVBAMMOE5ldXN0YXIgVUFUIEVudGVycHJpc2UgQ2VydGlmaWVkIENhbGxlciBJbnRlcm1lZGlhdGUgQ0ExMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEsmze7HXh1LddfmTnHN%2FaQyK9GPpn4PgVOJBa3sh40PETIJo3wAgKqvsdS3YM78yDiAZm%2BHlWuSmjJBq8dQ5RX6OB4TCB3jAPBgNVHRMBAf8EBTADAQH%2FMB8GA1UdIwQYMBaAFEsOaFjDkD3WaNSCONA2u933se6VMBUGA1UdIAQOMAwwCgYIKwYBBAHqAGQwZAYDVR0fBF0wWzBZoFegVYZTaHR0cHM6Ly9jcmwtdWF0LWVjY2lkLmNjaWQubmV1c3Rhci5iaXovTmV1c3RhckVudGVycHJpc2VDZXJ0aWZpZWRDYWxsZXJJZFJvb3RDQS5jcmwwHQYDVR0OBBYEFCt9OLm%2Fe3nm%2F9WmsNzYwh%2FjlrGDMA4GA1UdDwEB%2FwQEAwIBhjAKBggqhkjOPQQDAgNIADBFAiBEo6p8sRECsnfiLa9rsnnWbYo9Oy8Ra8Ukd7doJVOXbAIhAPkzxMozNRbydWrSlgGevqacrT7fOtzE48Cp5O909dDY)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_shaken_certificate_policies_id_ca](../../ISSUES/e_shaken_certificate_policies_id_ca/README.md) | error | US_SHAKEN_CP | The Certificate Policies extension contains an invalid OID value: 1.3.6.1.4.1.13568.100. Available OIDs: 2.16.840.1.114569.1.1.3, 2.16.840.1.114569.1.1.4 |
| [e_atis_subject_cn_ca](../../ISSUES/e_atis_subject_cn_ca/README.md) | error | ATIS1000080 | Common Name attribute 'Neustar UAT Enterprise Certified Caller Intermediate CA1' does not contain 'SHAKEN' |

### Not Effective

- e_atis_ext_not_specified_ca
- e_atis_serial_number_size_ca
- e_atis_subject_c_iso_ca

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 27 Nov 23 23:28 UTC