# Consumer module

The consumer module serves as an illustration of utilizing price information obtained from the pricefeed module.

### This is example of how to use pricefeed module in other module
```go
func (k Keeper) Price(c context.Context, req *types.QueryPriceRequest) (*types.QueryPriceResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	pf, err := k.priceFeedKeeper.GetPrice(ctx, req.Symbol)
	if err != nil {
		return nil, err
	}

	return &types.QueryPriceResponse{
		Symbol:      pf.Symbol,
		Price:       pf.Price,
		ResolveTime: pf.ResolveTime,
	}, nil
}
```
### Query latest price that got from BandChain

Query latest price that got from BandChin via consumer module

```
oracle-consumerd query consumer price BTC
```
