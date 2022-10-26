# STIR/SHAKEN CA Ecosystem Compliance

[Approved Certificate Authorities](https://authenticate.iconectiv.com/approved-certification-authorities) in the STIR/SHAKEN ecosystem are required to meet technical requirements from [ATIS-1000080](https://access.atis.org/apps/group_public/document.php?document_id=62163) and policy requirements from the supporting CA ecosystem’s [Certificate Policy](https://authenticate.iconectiv.com/documents-authenticate).

This report is generated using [Zlint](https://github.com/zmap/zlint) a tool commonly used to asses CA ecosystem compliance with such requirements. The tests used to generate this report are currently not part of the main Zlint distribution but can be found [here](https://github.com/martinisecurity/zlint).

## Summary

### Leaf Certificates

- Average validity span as of leaf certificates 364 days
- Percentage of leaf certificates expiring in the next 30 days is 38.40%

| Issuers | Certificates | Errors | Warnings | Notices | Not Effective |
|---------|--------------|--------|----------|---------|---------------|
| [Comcast](Comcast%2FREADME.md#leaf-certificates) | 37 (12.05%) | 37 (100.00%) | 37 (100.00%) | 0 (0.00%) | 0 (0.00%) |
| [GBSDTech](GBSDTech%2FREADME.md#leaf-certificates) | 1 (0.33%) | 1 (100.00%) | 1 (100.00%) | 0 (0.00%) | 0 (0.00%) |
| [Martini Security](Martini%20Security%2FREADME.md#leaf-certificates) | 7 (2.28%) | 7 (100.00%) | 0 (0.00%) | 4 (57.14%) | 0 (0.00%) |
| [Metaswitch](Metaswitch%2FREADME.md#leaf-certificates) | 21 (6.84%) | 1 (4.76%) | 0 (0.00%) | 0 (0.00%) | 20 (95.24%) |
| [NetNumber](NetNumber%2FREADME.md#leaf-certificates) | 4 (1.30%) | 4 (100.00%) | 1 (25.00%) | 4 (100.00%) | 0 (0.00%) |
| [Neustar](Neustar%2FREADME.md#leaf-certificates) | 97 (31.60%) | 65 (67.01%) | 4 (4.12%) | 0 (0.00%) | 33 (34.02%) |
| [Peeringhub](Peeringhub%2FREADME.md#leaf-certificates) | 5 (1.63%) | 5 (100.00%) | 5 (100.00%) | 0 (0.00%) | 0 (0.00%) |
| [Ribbon Communications](Ribbon%20Communications%2FREADME.md#leaf-certificates) | 3 (0.98%) | 3 (100.00%) | 3 (100.00%) | 0 (0.00%) | 0 (0.00%) |
| [Sansay](Sansay%2FREADME.md#leaf-certificates) | 75 (24.43%) | 75 (100.00%) | 75 (100.00%) | 0 (0.00%) | 0 (0.00%) |
| [T-Mobile](T-Mobile%2FREADME.md#leaf-certificates) | 1 (0.33%) | 1 (100.00%) | 1 (100.00%) | 0 (0.00%) | 0 (0.00%) |
| [TransNexus](TransNexus%2FREADME.md#leaf-certificates) | 56 (18.24%) | 56 (100.00%) | 55 (98.21%) | 0 (0.00%) | 2 (3.57%) |
| **Total** | 307 (100%) | 255 (83.06%) | 182 (59.28%) | 8 (2.61%) | 55 (17.92%) |

\* The percent of certificates per issuer is calculated against total certificates from all issuers.\
\*\* The percent of errors, warnings and notices is calculated against total observed certificates from the specified issuer.\
\*\*\* Tests use the ATIS 1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.\
\*\*\*\* For information on failed certificate repository retrievals see this [report](URL.md).\
\*\*\*\*\* 264 certificates skipped because they are currently expired.

### CA Certificates


| Issuers | Certificates | Errors | Warnings | Notices | Not Effective |
|---------|--------------|--------|----------|---------|---------------|
| [Comcast](Comcast%2FREADME.md#ca-certificates) | 2 (5.71%) | 0 (0.00%) | 0 (0.00%) | 1 (50.00%) | 2 (100.00%) |
| [GBSDTech](GBSDTech%2FREADME.md#ca-certificates) | 3 (8.57%) | 1 (33.33%) | 0 (0.00%) | 0 (0.00%) | 3 (100.00%) |
| [Martini Security](Martini%20Security%2FREADME.md#ca-certificates) | 3 (8.57%) | 2 (66.67%) | 1 (33.33%) | 1 (33.33%) | 2 (66.67%) |
| [Metaswitch](Metaswitch%2FREADME.md#ca-certificates) | 2 (5.71%) | 0 (0.00%) | 0 (0.00%) | 0 (0.00%) | 2 (100.00%) |
| [NetNumber](NetNumber%2FREADME.md#ca-certificates) | 3 (8.57%) | 0 (0.00%) | 0 (0.00%) | 0 (0.00%) | 3 (100.00%) |
| [Neustar](Neustar%2FREADME.md#ca-certificates) | 10 (28.57%) | 2 (20.00%) | 2 (20.00%) | 10 (100.00%) | 8 (80.00%) |
| [Peeringhub](Peeringhub%2FREADME.md#ca-certificates) | 2 (5.71%) | 1 (50.00%) | 0 (0.00%) | 2 (100.00%) | 2 (100.00%) |
| [Ribbon Communications](Ribbon%20Communications%2FREADME.md#ca-certificates) | 2 (5.71%) | 0 (0.00%) | 0 (0.00%) | 2 (100.00%) | 2 (100.00%) |
| [Sansay](Sansay%2FREADME.md#ca-certificates) | 2 (5.71%) | 1 (50.00%) | 1 (50.00%) | 1 (50.00%) | 1 (50.00%) |
| [T-Mobile](T-Mobile%2FREADME.md#ca-certificates) | 2 (5.71%) | 0 (0.00%) | 0 (0.00%) | 2 (100.00%) | 2 (100.00%) |
| [TransNexus](TransNexus%2FREADME.md#ca-certificates) | 4 (11.43%) | 3 (75.00%) | 0 (0.00%) | 4 (100.00%) | 4 (100.00%) |
| **Total** | 35 (100%) | 10 (28.57%) | 4 (11.43%) | 23 (65.71%) | 31 (88.57%) |

\* The percent of certificates per issuer is calculated against total certificates from all issuers.\
\*\* The percent of errors, warnings and notices is calculated against total observed certificates from the specified issuer.\
\*\*\* 2 certificates skipped because they are currently expired.

## Key

| Type | Description |
|------|-------------|
| Error | Tests in which the specifications are unambiguous on what the expected behavior must be. |
| Warning	| Tests in which the specifications are ambiguous or are provide only a recommendation. |
| Notice | Tests in which industry best practices are not followed. |
| Not Effective	| Tests that exist in the current specifications but were not in effect at the time of issuance. |

Generated: 26/10/2022 at 21:01:13