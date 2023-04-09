package api

import (
	"database/sql"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	mockdb "github.com/rezyfr/Trackerr-BackEnd/db/mock"
	db "github.com/rezyfr/Trackerr-BackEnd/db/sqlc"
	"github.com/rezyfr/Trackerr-BackEnd/token"
	"github.com/rezyfr/Trackerr-BackEnd/util"
	"github.com/stretchr/testify/require"
)

func TestListTransactions(t *testing.T) {
	user, _ := randomUser(t)
	n := 5
	transactions := make([]db.Transaction, n)
	for i := 0; i < n; i++ {
		transactions[i] = randomTransaction(user.ID)
	}

	type Query struct {
		pageLimit  int
		pageOffset int
	}

	testCases := []struct {
		query         Query
		name          string
		setupAuth     func(t *testing.T, request *http.Request, tokenMaker token.Maker)
		userId        int64
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:  "OK",
			query: Query{pageLimit: n, pageOffset: 1},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Email, int(user.ID), time.Minute)
			},
			userId: user.ID,
			buildStubs: func(store *mockdb.MockStore) {
				arg := db.ListTransactionsParams{
					Limit:  int32(n),
					Offset: 0,
					UserID: user.ID,
				}
				store.EXPECT().
					ListTransactions(gomock.Any(), gomock.Eq(arg)).
					Times(1).
					Return(transactions, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name:  "InternalError",
			query: Query{pageLimit: n, pageOffset: 1},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Email, int(user.ID), time.Minute)
			},
			userId: user.ID,
			buildStubs: func(store *mockdb.MockStore) {
				arg := db.ListTransactionsParams{
					Limit:  int32(n),
					Offset: 0,
					UserID: user.ID,
				}
				store.EXPECT().
					ListTransactions(gomock.Any(), gomock.Eq(arg)).
					Times(1).
					Return([]db.Transaction{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "InvalidLimit",
			query: Query{
				pageLimit:  51,
				pageOffset: 1,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Email, int(user.ID), time.Minute)
			},
			userId: user.ID,
			buildStubs: func(store *mockdb.MockStore) {
				arg := db.ListTransactionsParams{
					Limit:  int32(51),
					Offset: 0,
					UserID: user.ID,
				}
				store.EXPECT().
					ListTransactions(gomock.Any(), gomock.Eq(arg)).
					Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name:  "Unauthorized",
			query: Query{pageLimit: n, pageOffset: 1},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, "unauthorized@gmail.com", int(user.ID), time.Minute)
			},
			userId: user.ID,
			buildStubs: func(store *mockdb.MockStore) {
				arg := db.ListTransactionsParams{
					Limit:  int32(n),
					Offset: 0,
					UserID: user.ID,
				}
				store.EXPECT().
					ListTransactions(gomock.Any(), gomock.Eq(arg)).
					Times(1).
					Return(transactions, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		// {
		// 	name:  "No Authorization",
		// 	query: Query{pageLimit: n, pageOffset: 1},
		// 	setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
		// 		addAuthorization(t, request, tokenMaker, authorizationTypeBearer, "", int(user.ID), time.Minute)
		// 	},
		// 	userId: user.ID,
		// 	buildStubs: func(store *mockdb.MockStore) {
		// 		store.EXPECT().
		// 			ListTransactions(gomock.Any(), gomock.Any()).
		// 			Times(0)
		// 	},
		// 	checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
		// 		require.Equal(t, http.StatusUnauthorized, recorder.Code)
		// 	},
		// },
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			// start test server and send request
			server := newTestServer(t, store)
			recorder := httptest.NewRecorder()

			url := "/transactions"
			request, err := http.NewRequest(http.MethodGet, url, nil)

			q := request.URL.Query()
			q.Add("user_id", fmt.Sprintf("%d", tc.userId))
			q.Add("page_limit", fmt.Sprintf("%d", tc.query.pageLimit))
			q.Add("page_offset", fmt.Sprintf("%d", tc.query.pageOffset))
			request.URL.RawQuery = q.Encode()
			require.NoError(t, err)

			tc.setupAuth(t, request, server.tokenMaker)
			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}

func randomTransaction(userId int64) db.Transaction {
	return db.Transaction{
		ID:         util.RandomInt(1, 1000),
		UserID:     userId,
		Amount:     util.RandomInt(10000, 100000),
		Type:       db.Transactiontype(util.RandomType()),
		CategoryID: util.RandomInt(1, 1000),
		WalletID:   util.RandomInt(1, 1000),
	}
}
