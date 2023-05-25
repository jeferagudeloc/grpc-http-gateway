import { Container, Typography } from "@mui/material"


export const ProjectDescription = () => (
    <Container disableGutters maxWidth="lg" component="main" sx={{ pt: 8, pb: 6 }}>
    <Typography
        component="h1"
        variant="h2"
        align="center"
        color="text.primary"
        gutterBottom
      >
        gRPC HTTP Gateway
      </Typography>
      <Typography variant="h5" align="center" color="text.secondary" component="p">
      This proof of concept aims to compare the performance and velocity differences between gRPC and HTTP. 
      By conducting a series of tests and measurements, we will demonstrate the advantages and potential 
      speed improvements of using gRPC over traditional HTTP in certain scenarios.
      </Typography>
    </Container>
)