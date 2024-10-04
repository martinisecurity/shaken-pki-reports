# STIR/SHAKEN CA Ecosystem Compliance

[Approved Certificate Authorities](https://ecosystemcompliance.martinisecurity.com/#:~:text=Approved%20Certificate%20Authorities) in the STIR/SHAKEN ecosystem are required to meet technical requirements from [ATIS-1000080](https://access.atis.org/apps/group_public/document.php?document_id=62163) and policy requirements from the supporting CA ecosystem’s [Certificate Policy](https://authenticate.iconectiv.com/documents-authenticate).

This report is broken int two parts:
1. One generated using [Zlint](https://github.com/zmap/zlint) a tool commonly used to asses CA ecosystem compliance with such requirements. The tests used to generate this report are currently not part of the main [Zlint](https://github.com/martinisecurity/zlint) distribution but can be found here.
2. One generated with a custom script that eumerates the known STIR/SHAKEN certificates and asses each repository against the current rule set . The source for this test can be found here while the report itself can be found [here](REPOS/README.md).

## Summary

### Leaf Certificates

- 14684 certificates were included in the corpus being tested
- 1529 certificates in the corpus were skipped because they are duplicates
- 11652 certificates in the corpus were skipped because they are expired
- 472 certificates in the corpus were skipped because they are not currently trusted
- 1031 certificates being tested against the remaining rules
- 1.76 issues on average found in unexpired, trusted, and non-compliant certificates
- 49.56% of certificates contain one or more Error level issue
- 0.00% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 1.07% of certificates are too old to be assessed against currently enforced expectations
- 226 days is the average remaining validity for the certificates in the corpus
- 226 days is the average initial validity for the certificates in the corpus
- 408 certificates expire in the next 30 days
- 18.85 average number of unexpired certificates per OCN observed
- 779 unique OCNs observed in unexpired and valid certificate corpus

### CA Certificates

- 51 certificates were included in the corpus being tested
- 0 certificates in the corpus were skipped because they are duplicates
- 1 certificates in the corpus were skipped because they are expired
- 8 certificates in the corpus were skipped because they are not currently trusted
- 42 certificates being tested against the remaining rules
- 2.21 issues on average found in unexpired, trusted, and non-compliant certificates
- 45.24% of certificates contain one or more Error level issue
- 0.00% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 64.29% of certificates are too old to be assessed against currently enforced expectations
- 5836 days is the average remaining validity for the certificates in the corpus
- 5764 days is the average initial validity for the certificates in the corpus
- 0 certificates expire in the next 30 days

## Certificate Repository URL

- 49.85% of certificate repository URLs contain one or more Error level issue
- 51.31% of certificates repository URLs contain one or more Warning level issue
- 0.00% of certificates repository URLs contain one or more Notice level issue

## Details

\* The percent of certificates per issuer is calculated against total certificates from all issuers.\
\*\* The percent of errors, warnings and notices is calculated against total observed certificates from the specified issuer.\
\*\*\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.

### Leaf Certificates

| Issuers | Certificates | Errors | Warnings | Notices | Not Effective |
|---------|--------------|--------|----------|---------|---------------|
| [Comcast](CERTS/Comcast/README.md#leaf-certificates) | 31 (3.01%) | 31 (100.00%) | 0 (0.00%) | 0 (0.00%) | 0 (0.00%) |
| [GBSDTech](CERTS/GBSDTech/README.md#leaf-certificates) | 5 (0.48%) | 5 (100.00%) | 0 (0.00%) | 0 (0.00%) | 0 (0.00%) |
| [Martini Security](CERTS/Martini_Security/README.md#leaf-certificates) | 391 (37.92%) | 0 (0.00%) | 0 (0.00%) | 0 (0.00%) | 0 (0.00%) |
| [Metaswitch](CERTS/Metaswitch/README.md#leaf-certificates) | 65 (6.30%) | 65 (100.00%) | 0 (0.00%) | 0 (0.00%) | 11 (16.92%) |
| [NetNumber](CERTS/NetNumber/README.md#leaf-certificates) | 6 (0.58%) | 6 (100.00%) | 0 (0.00%) | 0 (0.00%) | 0 (0.00%) |
| [Neustar](CERTS/Neustar/README.md#leaf-certificates) | 144 (13.97%) | 144 (100.00%) | 0 (0.00%) | 0 (0.00%) | 0 (0.00%) |
| [Peeringhub](CERTS/Peeringhub/README.md#leaf-certificates) | 30 (2.91%) | 1 (3.33%) | 0 (0.00%) | 0 (0.00%) | 0 (0.00%) |
| [Ribbon Communications](CERTS/Ribbon_Communications/README.md#leaf-certificates) | 14 (1.36%) | 14 (100.00%) | 0 (0.00%) | 0 (0.00%) | 0 (0.00%) |
| [Sansay](CERTS/Sansay/README.md#leaf-certificates) | 233 (22.60%) | 233 (100.00%) | 0 (0.00%) | 0 (0.00%) | 0 (0.00%) |
| [T-Mobile](CERTS/T-Mobile/README.md#leaf-certificates) | 1 (0.10%) | 1 (100.00%) | 0 (0.00%) | 0 (0.00%) | 0 (0.00%) |
| [Telonium](CERTS/Telonium/README.md#leaf-certificates) | 25 (2.42%) | 6 (24.00%) | 0 (0.00%) | 0 (0.00%) | 0 (0.00%) |
| [TransNexus](CERTS/TransNexus/README.md#leaf-certificates) | 86 (8.34%) | 5 (5.81%) | 0 (0.00%) | 0 (0.00%) | 0 (0.00%) |
| **Total** | 1031 (100.00%) | 511 (49.56%) | 0 (0.00%) | 0 (0.00%) | 11 (1.07%) |

### CA Certificates

| Issuers | Certificates | Errors | Warnings | Notices | Not Effective |
|---------|--------------|--------|----------|---------|---------------|
| [CTIA](CERTS/CTIA/README.md#ca-certificates) | 1 (2.38%) | 1 (100.00%) | 0 (0.00%) | 0 (0.00%) | 0 (0.00%) |
| [Comcast](CERTS/Comcast/README.md#ca-certificates) | 2 (4.76%) | 1 (50.00%) | 0 (0.00%) | 0 (0.00%) | 1 (50.00%) |
| [GBSDTech](CERTS/GBSDTech/README.md#ca-certificates) | 3 (7.14%) | 2 (66.67%) | 0 (0.00%) | 0 (0.00%) | 2 (66.67%) |
| [Martini Security](CERTS/Martini_Security/README.md#ca-certificates) | 4 (9.52%) | 0 (0.00%) | 0 (0.00%) | 0 (0.00%) | 2 (50.00%) |
| [Metaswitch](CERTS/Metaswitch/README.md#ca-certificates) | 4 (9.52%) | 4 (100.00%) | 0 (0.00%) | 0 (0.00%) | 2 (50.00%) |
| [NetNumber](CERTS/NetNumber/README.md#ca-certificates) | 3 (7.14%) | 2 (66.67%) | 0 (0.00%) | 0 (0.00%) | 3 (100.00%) |
| [Neustar](CERTS/Neustar/README.md#ca-certificates) | 6 (14.29%) | 1 (16.67%) | 0 (0.00%) | 0 (0.00%) | 6 (100.00%) |
| [Peeringhub](CERTS/Peeringhub/README.md#ca-certificates) | 2 (4.76%) | 2 (100.00%) | 0 (0.00%) | 0 (0.00%) | 2 (100.00%) |
| [Ribbon Communications](CERTS/Ribbon_Communications/README.md#ca-certificates) | 3 (7.14%) | 2 (66.67%) | 0 (0.00%) | 0 (0.00%) | 2 (66.67%) |
| [SOMOS](CERTS/SOMOS/README.md#ca-certificates) | 1 (2.38%) | 0 (0.00%) | 0 (0.00%) | 0 (0.00%) | 0 (0.00%) |
| [Sansay](CERTS/Sansay/README.md#ca-certificates) | 2 (4.76%) | 0 (0.00%) | 0 (0.00%) | 0 (0.00%) | 2 (100.00%) |
| [T-Mobile](CERTS/T-Mobile/README.md#ca-certificates) | 3 (7.14%) | 1 (33.33%) | 0 (0.00%) | 0 (0.00%) | 3 (100.00%) |
| [Telonium](CERTS/Telonium/README.md#ca-certificates) | 5 (11.90%) | 3 (60.00%) | 0 (0.00%) | 0 (0.00%) | 0 (0.00%) |
| [TransNexus](CERTS/TransNexus/README.md#ca-certificates) | 3 (7.14%) | 0 (0.00%) | 0 (0.00%) | 0 (0.00%) | 2 (66.67%) |
| **Total** | 42 (100.00%) | 19 (45.24%) | 0 (0.00%) | 0 (0.00%) | 27 (64.29%) |

### Key

| Type | Description |
|------|-------------|
| Errors | Tests in which the specifications are unambiguous on what the expected behavior must be. |
| Warnings | Tests in which the specifications are ambiguous or are provide only a recommendation. |
| Notices | Tests in which industry best practices are not followed. |
| Not Effective | Tests that exist in the current specifications but were not in effect at the time of issuance. |


Generated: 04 Oct 24 16:29 UTC