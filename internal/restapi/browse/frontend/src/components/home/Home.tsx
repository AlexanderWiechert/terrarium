import Header from "../header/Header";
import {Container} from "@mui/material";
import {Outlet} from "react-router-dom";
import Footer from "../footer/Footer";
import React from "react";

const Home = () => {
    return <>
        <Header />
        <Container style={{marginTop: ".3em", marginBottom: ".85em"}}>
            <Outlet />
        </Container>
        <Footer />
    </>
}
export default Home

