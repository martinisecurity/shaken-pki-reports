# STIR/SHAKEN CA Ecosystem Compliance

[Approved Certificate Authorities](https://ecosystemcompliance.martinisecurity.com/#:~:text=Approved%20Certificate%20Authorities) in the STIR/SHAKEN ecosystem are required to meet technical requirements from [ATIS-1000080](https://access.atis.org/apps/group_public/document.php?document_id=62163) and policy requirements from the supporting CA ecosystem’s [Certificate Policy](https://authenticate.iconectiv.com/documents-authenticate).

This report is broken int two parts:
1. One generated using [Zlint](https://github.com/zmap/zlint) a tool commonly used to asses CA ecosystem compliance with such requirements. The tests used to generate this report are currently not part of the main [Zlint](https://github.com/martinisecurity/zlint) distribution but can be found here.
2. One generated with a custom script that eumerates the known STIR/SHAKEN certificates and asses each repository against the current rule set . The source for this test can be found here while the report itself can be found [here](REPOS/README.md).

## Summary

### Leaf Certificates

- 3149 certificates were included in the corpus being tested
- 224 certificates in the corpus were skipped because they are duplicates
- 2483 certificates in the corpus were skipped because they are expired
- 9 certificates in the corpus were skipped because they are not currently trusted
- 433 certificates being tested against the remaining rules
- 3.89 issues on average found in unexpired, trusted, and non-compliant certificates
- 90.76% of certificates contain one or more Error level issue
- 50.81% of certificates contain one or more Warning level issue
- 1.39% of certificates contain one or more Notice level issue
- 14.78% of certificates are too old to be assessed against currently enforced expectations
- 341 days is the average remaining validity for the certificates in the corpus
- 342 days is the average initial validity for the certificates in the corpus
- 174 certificates expire in the next 30 days
- 7.50 average number of unexpired certificates per OCN observed
- 420 unique OCNs observed in unexpired and valid certificate corpus

### CA Certificates

- 40 certificates were included in the corpus being tested
- 0 certificates in the corpus were skipped because they are duplicates
- 1 certificates in the corpus were skipped because they are expired
- 4 certificates in the corpus were skipped because they are not currently trusted
- 35 certificates being tested against the remaining rules
- 2.04 issues on average found in unexpired, trusted, and non-compliant certificates
- 54.29% of certificates contain one or more Error level issue
- 48.57% of certificates contain one or more Warning level issue
- 2.86% of certificates contain one or more Notice level issue
- 65.71% of certificates are too old to be assessed against currently enforced expectations
- 5569 days is the average remaining validity for the certificates in the corpus
- 5573 days is the average initial validity for the certificates in the corpus
- 0 certificates expire in the next 30 days

## Certificate Repository URL

- 52.89% of certificate repository URLs contain one or more Error level issue
- 85.22% of certificates repository URLs contain one or more Warning level issue
- 0.00% of certificates repository URLs contain one or more Notice level issue

## Details

\* The percent of certificates per issuer is calculated against total certificates from all issuers.\
\*\* The percent of errors, warnings and notices is calculated against total observed certificates from the specified issuer.\
\*\*\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.

### Leaf Certificates

| Issuers | Certificates | Errors | Warnings | Notices | Not Effective |
|---------|--------------|--------|----------|---------|---------------|
| [Comcast](CERTS/Comcast/README.md#leaf-certificates) | 20 (4.62%) | 20 (100.00%) | 20 (100.00%) | 0 (0.00%) | 0 (0.00%) |
| [GBSDTech](CERTS/GBSDTech/README.md#leaf-certificates) | 3 (0.69%) | 3 (100.00%) | 0 (0.00%) | 0 (0.00%) | 0 (0.00%) |
| [Martini Security](CERTS/Martini_Security/README.md#leaf-certificates) | 40 (9.24%) | 0 (0.00%) | 0 (0.00%) | 1 (2.50%) | 0 (0.00%) |
| [Metaswitch](CERTS/Metaswitch/README.md#leaf-certificates) | 43 (9.93%) | 43 (100.00%) | 1 (2.33%) | 0 (0.00%) | 37 (86.05%) |
| [NetNumber](CERTS/NetNumber/README.md#leaf-certificates) | 5 (1.15%) | 5 (100.00%) | 0 (0.00%) | 5 (100.00%) | 0 (0.00%) |
| [Neustar](CERTS/Neustar/README.md#leaf-certificates) | 127 (29.33%) | 127 (100.00%) | 29 (22.83%) | 0 (0.00%) | 27 (21.26%) |
| [Peeringhub](CERTS/Peeringhub/README.md#leaf-certificates) | 9 (2.08%) | 9 (100.00%) | 1 (11.11%) | 0 (0.00%) | 0 (0.00%) |
| [Ribbon Communications](CERTS/Ribbon_Communications/README.md#leaf-certificates) | 8 (1.85%) | 8 (100.00%) | 8 (100.00%) | 0 (0.00%) | 0 (0.00%) |
| [Sansay](CERTS/Sansay/README.md#leaf-certificates) | 151 (34.87%) | 151 (100.00%) | 151 (100.00%) | 0 (0.00%) | 0 (0.00%) |
| [T-Mobile](CERTS/T-Mobile/README.md#leaf-certificates) | 2 (0.46%) | 2 (100.00%) | 2 (100.00%) | 0 (0.00%) | 0 (0.00%) |
| [TransNexus](CERTS/TransNexus/README.md#leaf-certificates) | 25 (5.77%) | 25 (100.00%) | 8 (32.00%) | 0 (0.00%) | 0 (0.00%) |
| **Total** | 433 (100.00%) | 393 (90.76%) | 220 (50.81%) | 6 (1.39%) | 64 (14.78%) |

### CA Certificates

| Issuers | Certificates | Errors | Warnings | Notices | Not Effective |
|---------|--------------|--------|----------|---------|---------------|
| [Comcast](CERTS/Comcast/README.md#ca-certificates) | 2 (5.71%) | 0 (0.00%) | 0 (0.00%) | 0 (0.00%) | 2 (100.00%) |
| [GBSDTech](CERTS/GBSDTech/README.md#ca-certificates) | 2 (5.71%) | 1 (50.00%) | 0 (0.00%) | 0 (0.00%) | 2 (100.00%) |
| [Martini Security](CERTS/Martini_Security/README.md#ca-certificates) | 3 (8.57%) | 0 (0.00%) | 0 (0.00%) | 1 (33.33%) | 0 (0.00%) |
| [Metaswitch](CERTS/Metaswitch/README.md#ca-certificates) | 2 (5.71%) | 2 (100.00%) | 0 (0.00%) | 0 (0.00%) | 2 (100.00%) |
| [NetNumber](CERTS/NetNumber/README.md#ca-certificates) | 3 (8.57%) | 2 (66.67%) | 2 (66.67%) | 0 (0.00%) | 3 (100.00%) |
| [Neustar](CERTS/Neustar/README.md#ca-certificates) | 6 (17.14%) | 4 (66.67%) | 6 (100.00%) | 0 (0.00%) | 4 (66.67%) |
| [Peeringhub](CERTS/Peeringhub/README.md#ca-certificates) | 2 (5.71%) | 2 (100.00%) | 2 (100.00%) | 0 (0.00%) | 1 (50.00%) |
| [Ribbon Communications](CERTS/Ribbon_Communications/README.md#ca-certificates) | 2 (5.71%) | 2 (100.00%) | 2 (100.00%) | 0 (0.00%) | 2 (100.00%) |
| [Sansay](CERTS/Sansay/README.md#ca-certificates) | 2 (5.71%) | 0 (0.00%) | 2 (100.00%) | 0 (0.00%) | 1 (50.00%) |
| [T-Mobile](CERTS/T-Mobile/README.md#ca-certificates) | 4 (11.43%) | 1 (25.00%) | 0 (0.00%) | 0 (0.00%) | 3 (75.00%) |
| [Telonium](CERTS/Telonium/README.md#ca-certificates) | 2 (5.71%) | 2 (100.00%) | 0 (0.00%) | 0 (0.00%) | 0 (0.00%) |
| [TransNexus](CERTS/TransNexus/README.md#ca-certificates) | 5 (14.29%) | 3 (60.00%) | 3 (60.00%) | 0 (0.00%) | 3 (60.00%) |
| **Total** | 35 (100.00%) | 19 (54.29%) | 17 (48.57%) | 1 (2.86%) | 23 (65.71%) |

### Key

| Type | Description |
|------|-------------|
| Errors | Tests in which the specifications are unambiguous on what the expected behavior must be. |
| Warnings | Tests in which the specifications are ambiguous or are provide only a recommendation. |
| Notices | Tests in which industry best practices are not followed. |
| Not Effective | Tests that exist in the current specifications but were not in effect at the time of issuance. |


Generated: 12 Apr 23 01:46 UTC