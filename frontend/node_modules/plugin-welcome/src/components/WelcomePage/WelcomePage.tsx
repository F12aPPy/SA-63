import React, { FC } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import ComponanceTable from '../Table';
import Button from '@material-ui/core/Button';
 
import {
 Content,
 Header,
 Page,
 pageTheme,
 ContentHeader,
 Link,
} from '@backstage/core';
 
const WelcomePage: FC<{}> = () => {
 const profile = { givenName: 'ระบบฝากครรภ์' };
 
 return (
   <Page theme={pageTheme.home}>
     <Header
       title={`Welcome to ${profile.givenName || 'to Backstage'}`}
       subtitle="นพ.AAA BBBB"
     ></Header>
     <Content>
       <ContentHeader title="ข้อมูลระบบฝากครรภ์">
         <Link component={RouterLink} to="/user">
           <Button variant="contained" color="primary">
             Add User
           </Button>
         </Link>
       </ContentHeader>
       <ComponanceTable></ComponanceTable>
     </Content>
   </Page>
 );
};
 
export default WelcomePage;
