import { useState, useEffect } from 'react';
import { showError, renderQuota } from 'utils/common';

import Table from '@mui/material/Table';
import TableBody from '@mui/material/TableBody';
import TableContainer from '@mui/material/TableContainer';
import PerfectScrollbar from 'react-perfect-scrollbar';
import TablePagination from '@mui/material/TablePagination';
import LinearProgress from '@mui/material/LinearProgress';
import ButtonGroup from '@mui/material/ButtonGroup';
import Toolbar from '@mui/material/Toolbar';

import { Button, Card, Stack, Container, Typography, Box, Alert, IconButton, Chip } from '@mui/material';
import { IconRefresh, IconSearch } from '@tabler/icons-react';
import LogTableRow from './component/TableRow';
import LogTableHead from './component/TableHead';
import TableToolBar from './component/TableToolBar';
import { API } from 'utils/api';
import { isAdmin } from 'utils/common';
import { ITEMS_PER_PAGE } from 'constants';

export default function Log() {
  const originalKeyword = {
    p: 0,
    username: '',
    token_name: '',
    model_name: '',
    start_timestamp: 0,
    end_timestamp: new Date().getTime() / 1000 + 3600,
    type: 0,
    channel: ''
  };
  const [logs, setLogs] = useState([]);
  const [activePage, setActivePage] = useState(0);
  const [searching, setSearching] = useState(false);
  const [searchKeyword, setSearchKeyword] = useState(originalKeyword);
  const [initPage, setInitPage] = useState(true);
  const [stat, setStat] = useState({ quota: 0 });
  const [showStat, setShowStat] = useState(false);
  const [isStatRefreshing, setIsStatRefreshing] = useState(false);
  const userIsAdmin = isAdmin();

  const loadLogs = async (startIdx) => {
    setSearching(true);
    const url = userIsAdmin ? '/api/log/' : '/api/log/self';
    const query = searchKeyword;

    query.p = startIdx;
    if (!userIsAdmin) {
      delete query.username;
      delete query.channel;
    }
    const res = await API.get(url, { params: query });
    const { success, message, data } = res.data;
    if (success) {
      if (startIdx === 0) {
        setLogs(data);
      } else {
        let newLogs = [...logs];
        newLogs.splice(startIdx * ITEMS_PER_PAGE, data.length, ...data);
        setLogs(newLogs);
      }
    } else {
      showError(message);
    }
    setSearching(false);
  };

  const onPaginationChange = (event, activePage) => {
    (async () => {
      if (activePage === Math.ceil(logs.length / ITEMS_PER_PAGE)) {
        // In this case we have to load more data and then append them.
        await loadLogs(activePage);
      }
      setActivePage(activePage);
    })();
  };

  const searchLogs = async (event) => {
    event.preventDefault();
    await loadLogs(0);
    setActivePage(0);
    return;
  };

  const handleSearchKeyword = (event) => {
    setSearchKeyword({ ...searchKeyword, [event.target.name]: event.target.value });
  };

  const getLogStat = async () => {
    const query = { ...searchKeyword };
    delete query.p;
    if (!userIsAdmin) {
      delete query.username;
      delete query.channel;
    }
    const url = userIsAdmin ? '/api/log/stat' : '/api/log/self/stat';
    const res = await API.get(url, { params: query });
    const { success, message, data } = res.data;
    if (success) {
      setStat(data);
    } else {
      showError(message);
    }
  };

  const handleStatRefresh = async () => {
    setIsStatRefreshing(true);
    try {
      await getLogStat();
    } finally {
      setIsStatRefreshing(false);
    }
  };

  const handleShowStat = async () => {
    if (!showStat) {
      await getLogStat();
    }
    setShowStat(!showStat);
  };

  // 处理刷新
  const handleRefresh = () => {
    setInitPage(true);
  };

  useEffect(() => {
    setSearchKeyword(originalKeyword);
    setActivePage(0);
    loadLogs(0)
      .then()
      .catch((reason) => {
        showError(reason);
      });
    setInitPage(false);
  }, [initPage]);

  return (
    <>
      <Stack direction="row" alignItems="center" justifyContent="space-between" mb={2.5}>
        <Typography variant="h4">日志</Typography>
      </Stack>

      {/* Quota Statistics Section */}
      <Card sx={{ mb: 2 }}>
        <Box sx={{ p: 2 }}>
          <Stack direction="row" alignItems="center" spacing={2}>
            <Typography variant="h6">
              使用详情（总配额：
              {showStat ? (
                <>
                  <Chip
                    label={renderQuota(stat.quota)}
                    color="primary"
                    variant="outlined"
                    size="small"
                  />
                  <IconButton
                    size="small"
                    onClick={handleStatRefresh}
                    disabled={isStatRefreshing}
                    sx={{ ml: 1 }}
                    title="刷新配额数据"
                  >
                    <IconRefresh
                      width="16px"
                      style={{
                        animation: isStatRefreshing ? 'spin 1s linear infinite' : 'none'
                      }}
                    />
                  </IconButton>
                </>
              ) : (
                <Button
                  size="small"
                  variant="text"
                  onClick={handleShowStat}
                  sx={{ textTransform: 'none', color: 'text.secondary' }}
                >
                  点击查看
                </Button>
              )}
              ）
            </Typography>
          </Stack>
        </Box>
      </Card>
      <Card>
        <Box component="form" onSubmit={searchLogs} noValidate sx={{marginTop: 2}}>
          <TableToolBar filterName={searchKeyword} handleFilterName={handleSearchKeyword} userIsAdmin={userIsAdmin} />
        </Box>
        <Toolbar
          sx={{
            textAlign: 'right',
            height: 50,
            display: 'flex',
            justifyContent: 'space-between',
            p: (theme) => theme.spacing(0, 1, 0, 3)
          }}
        >
          <Container>
            <ButtonGroup variant="outlined" aria-label="outlined small primary button group" sx={{marginBottom: 2}}>
              <Button onClick={handleRefresh} startIcon={<IconRefresh width={'18px'} />}>
                刷新/清除搜索条件
              </Button>

              <Button onClick={searchLogs} startIcon={<IconSearch width={'18px'} />}>
                搜索
              </Button>
            </ButtonGroup>
          </Container>
        </Toolbar>
        {searching && <LinearProgress />}
        <PerfectScrollbar component="div">
          <TableContainer sx={{ overflow: 'unset' }}>
            <Table sx={{ minWidth: 800 }}>
              <LogTableHead userIsAdmin={userIsAdmin} />
              <TableBody>
                {logs.slice(activePage * ITEMS_PER_PAGE, (activePage + 1) * ITEMS_PER_PAGE).map((row, index) => (
                  <LogTableRow item={row} key={`${row.id}_${index}`} userIsAdmin={userIsAdmin} />
                ))}
              </TableBody>
            </Table>
          </TableContainer>
        </PerfectScrollbar>
        <TablePagination
          page={activePage}
          component="div"
          count={logs.length + (logs.length % ITEMS_PER_PAGE === 0 ? 1 : 0)}
          rowsPerPage={ITEMS_PER_PAGE}
          onPageChange={onPaginationChange}
          rowsPerPageOptions={[ITEMS_PER_PAGE]}
        />
      </Card>
    </>
  );
}
