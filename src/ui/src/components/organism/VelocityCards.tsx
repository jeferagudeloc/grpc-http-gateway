import { Box, Button, Card, CardActions, CardContent, CardHeader, Container, Grid, Typography } from "@mui/material";
import { useGetOrdersGrpcQuery, useGetUsersGrpcQuery, useGetOrdersHttpQuery, useGetUsersHttpQuery } from '../../store/api/grpcGatewayApi';
import { useState } from 'react';
import ReplayIcon from '@mui/icons-material/Replay';
export interface Request {
  id: string;
  title: string;
  time?: string;
  description: string[];
  buttonText: string;
  buttonVariant: "outlined" | "contained";
}

interface VelocityCardsProps {
  requests: Request[];
}

export const VelocityCards = ({ requests: initialRequests }: VelocityCardsProps) => {
  const { refetch: getOrdersGrpcQueryRefetch } = useGetOrdersGrpcQuery();
  const { refetch: getUsersGrpcQueryRefetch } = useGetUsersGrpcQuery();
  const { refetch: getOrdersHttpQueryRefetch } = useGetOrdersHttpQuery();
  const { refetch: getUsersHttpQueryRefetch } = useGetUsersHttpQuery();

  const [requests, setRequests] = useState<Request[]>(initialRequests);

  const handleFetchRequest = async (refetch: () => Promise<any>, requestId: string) => {
    const startTime = new Date().getTime();
    refetch().then((response) => {
      const endTime = new Date().getTime();
      const duration = endTime - startTime;
      updateRequestTime(requestId, duration);
    }).catch(err => console.error(err));
    
  };

  const updateRequestTime = (requestId: string, duration: number) => {
    setRequests((prevRequests) =>
      prevRequests.map((request) =>
        request.id === requestId ? { ...request, time: `${duration} ms` } : request
      )
    );
  };

  const fetchAll = async (requestId?: string) => {
    const fetchers = [
      { refetch: getOrdersGrpcQueryRefetch, requestId: "gRPC-getOrders" },
      { refetch: getUsersGrpcQueryRefetch, requestId: "gRPC-getUsers" },
      { refetch: getOrdersHttpQueryRefetch, requestId: "HTTP-getOrders" },
      { refetch: getUsersHttpQueryRefetch, requestId: "HTTP-getUsers" },
    ];

    if (requestId) {
      const fetcher = fetchers.find((fetcher) => fetcher.requestId === requestId);
      if (fetcher) {
        await handleFetchRequest(fetcher.refetch, fetcher.requestId);
      }
    } else {
      setRequests(initialRequests)
      await Promise.all(fetchers.map((fetcher) => handleFetchRequest(fetcher.refetch, fetcher.requestId)));
    }
  };

  return (
    <Container maxWidth="xl">
      <Button fullWidth onClick={() => fetchAll()} variant="contained" endIcon={<ReplayIcon />}>
        Call all APIs
      </Button>
      <Grid container spacing={2} alignItems="center" paddingTop={5} justifyContent="center">
        {requests.map((request) => (
          <Grid item key={request.id} xs={12} md={3}>
            <Card>
              <CardHeader
                title={request.title}
                titleTypographyProps={{ align: "center" }}
                subheaderTypographyProps={{ align: "center" }}
                sx={{
                  backgroundColor: (theme) =>
                    theme.palette.mode === "light" ? theme.palette.grey[200] : theme.palette.grey[700],
                    borderBottom: '2px solid',
                    borderBottomImage: 'linear-gradient(-120deg, #4285f4, #34a853, #fbbc05, #ea4335)',
                }}
              />
              <CardContent>
                <Box
                  sx={{
                    display: "flex",
                    justifyContent: "center",
                    alignItems: "baseline",
                    mb: 2,
                  }}
                >
                  <Typography component="h2" variant="h3" color="text.primary">
                    { request.time || "-" }
                  </Typography>
                </Box>
                <ul>
                  {request.description.map((line) => (
                    <Typography component="li" variant="subtitle1" align="center" key={line}>
                      {line}
                    </Typography>
                  ))}
                </ul>
              </CardContent>
              <CardActions>
                <Button fullWidth variant={'outlined'} onClick={() => fetchAll(request.id)}>
                  {request.buttonText}
                </Button>
              </CardActions>
            </Card>
          </Grid>
        ))}
      </Grid>
    </Container>
  );
};