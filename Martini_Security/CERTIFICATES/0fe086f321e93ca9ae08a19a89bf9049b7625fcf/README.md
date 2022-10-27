# STIR/SHAKEN CA Ecosystem Compliance
## Martini Security

### Certificate 0fe086f321e93ca9ae08a19a89bf9049b7625fcf
Tested At: 2022-10-27 18:55:39 +0000 UTC\
Initial Validity Period: 1825 day(s)\
Remaining Validity Period: 1648 day(s)\
Subject: CN=Martini Security SHAKEN G1, O=Martini Security\\, LLC, L=Seattle, ST=WA, C=US\
Issuer: CN=Martini Security SHAKEN R1, O=Martini Security\\, LLC, L=Seattle, ST=WA, C=US

Link: https://p.mtsec.me/2dd5/HANi-8RbpIVe.pem

View: [Click to view](https://understandingwebpki.com/?cert=MIIDEjCCArmgAwIBAgIUWRP%2FeaQF%2B0VNFxze3w%2BczApA53owCgYIKoZIzj0EAwIwcTELMAkGA1UEBhMCVVMxCzAJBgNVBAgTAldBMRAwDgYDVQQHEwdTZWF0dGxlMR4wHAYDVQQKExVNYXJ0aW5pIFNlY3VyaXR5LCBMTEMxIzAhBgNVBAMTGk1hcnRpbmkgU2VjdXJpdHkgU0hBS0VOIFIxMB4XDTIyMDUwMzE1NDEwMFoXDTI3MDUwMjE1NDEwMFowcTELMAkGA1UEBhMCVVMxCzAJBgNVBAgTAldBMRAwDgYDVQQHEwdTZWF0dGxlMR4wHAYDVQQKExVNYXJ0aW5pIFNlY3VyaXR5LCBMTEMxIzAhBgNVBAMTGk1hcnRpbmkgU2VjdXJpdHkgU0hBS0VOIEcxMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEgHMKTCPENlBRYIP%2FLPOk528hT0mhODS%2BqQkHt5ujQkKSfgLeossZ1xb%2FvZFlUIIvz67C6Ugb5nupkadMzFQWZ6OCAS0wggEpMA4GA1UdDwEB%2FwQEAwICBDASBgNVHRMBAf8ECDAGAQH%2FAgEAMB0GA1UdDgQWBBQt1ctQCVNMkSR%2FI7Gp5SWHOJZJgjAfBgNVHSMEGDAWgBSLmgJLXKXv3jruIFZxYMWJNOBIeTCBpgYDVR0fBIGeMIGbMIGYoDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMRQwEgYDVQQHDAtCcmlkZ2V3YXRlcjELMAkGA1UECAwCTkoxEzARBgNVBAMMClNUSS1QQSBDUkwxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEEwGgYDVR0gAQH%2FBBAwDjAMBgpghkgBhv8JAQEDMAoGCCqGSM49BAMCA0cAMEQCICYdBOW7vfq5o01UgzJTmpu6qfjKbZA%2FthPQEFb68e%2BwAiAmSMR%2FRPxKHvbKGa4z8ktL5XUwkI2fbbbxe322Vke0HQ%3D%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| e_sti_ca_crl_distribution | error | ATIS-1000080v4 | CRL Distribution Point shall be reachable if the requesting IP address within the program ACLs |
| n_sti_ca_certificate_policy_critical | notice | ATIS-1000080v4 | STI certificates should contain a CertificatePolicies extension marked uncritical |

### Not Effective

- e_cp1_3_ca_key_usage_crl_sign
- w_cp1_3_ca_subject_rdn_unknown

\* Tests use the ATIS 1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.

Generated: 27/10/2022 at 18:57:26