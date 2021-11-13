import './game.module.less';
import React from 'react'
import MapComponent from '../map';
import { GameProvider } from "../provider"
import { Map as MapState, Params, State, Overview } from "../../codec/rook/game/game"
import Long from 'long';

export interface GameProps {
  gameID: Long
  provider: GameProvider
}

export interface GameState {
  overview: Overview,
  state: State,
  params: Params
  error?: Error
}

class Game extends React.Component<GameProps, GameState> {

  async componentDidMount() {
    await this.load(this.props.gameID)
  }

  async load(gameID: Long): Promise<void> {
    try { 
      const gameResp = await this.props.provider.query.FindByID({ id: gameID })
      const stateResp = await this.props.provider.query.State({ id: gameID })
      if (!gameResp.overview || !stateResp.state) {
        console.error("game not found")
        return
      }
      const params = await this.props.provider.query.Params({ version: gameResp.overview.paramVersion})
      if (!params.params) {
        console.error("params not found")
        return
      }
      this.setState({ 
        overview: gameResp.overview,
        state: stateResp.state,
        params: params.params,
      })
    } catch (err) {
      this.setState({ error: err as Error })
    }
  }

  render() {
    return (
      <MapComponent map={this.state.overview.map!}/>
    );
  }
}

export default Game;
