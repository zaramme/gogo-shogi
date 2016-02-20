package pieceScore

import (
	p "github.com/zaramme/gogo-shogi/piece"
	s "github.com/zaramme/gogo-shogi/score"
)

const (
	PawnScore      s.Score = 100 * 9 / 10
	LanceScore     s.Score = 350 * 9 / 10
	KnightScore    s.Score = 450 * 9 / 10
	SilverScore    s.Score = 550 * 9 / 10
	GoldScore      s.Score = 600 * 9 / 10
	BishopScore    s.Score = 950 * 9 / 10
	RookScore      s.Score = 1100 * 9 / 10
	ProPawnScore   s.Score = 600 * 9 / 10
	ProLanceScore  s.Score = 600 * 9 / 10
	ProKnightScore s.Score = 600 * 9 / 10
	ProSilverScore s.Score = 600 * 9 / 10
	HorseScore     s.Score = 1050 * 9 / 10
	DragonScore    s.Score = 1550 * 9 / 10

	KingScore s.Score = 15000

	CapturePawnScore      s.Score = PawnScore * 2
	CaptureLanceScore     s.Score = LanceScore * 2
	CaptureKnightScore    s.Score = KnightScore * 2
	CaptureSilverScore    s.Score = SilverScore * 2
	CaptureGoldScore      s.Score = GoldScore * 2
	CaptureBishopScore    s.Score = BishopScore * 2
	CaptureRookScore      s.Score = RookScore * 2
	CaptureProPawnScore   s.Score = ProPawnScore + PawnScore
	CaptureProLanceScore  s.Score = ProLanceScore + LanceScore
	CaptureProKnightScore s.Score = ProKnightScore + KnightScore
	CaptureProSilverScore s.Score = ProSilverScore + SilverScore
	CaptureHorseScore     s.Score = HorseScore + BishopScore
	CaptureDragonScore    s.Score = DragonScore + RookScore
	CaptureKingScore      s.Score = KingScore * 2

	PromotePawnScore   s.Score = ProPawnScore - PawnScore
	PromoteLanceScore  s.Score = ProLanceScore - LanceScore
	PromoteKnightScore s.Score = ProKnightScore - KnightScore
	PromoteSilverScore s.Score = ProSilverScore - SilverScore
	PromoteBishopScore s.Score = HorseScore - BishopScore
	PromoteRookScore   s.Score = DragonScore - RookScore

	ScoreKnownWin s.Score = KingScore
)

var PieceScore = [p.PieceNone]s.Score{
	s.Zero,
	PawnScore, LanceScore, KnightScore, SilverScore, BishopScore, RookScore, GoldScore,
	s.Zero, // King
	ProPawnScore, ProLanceScore, ProKnightScore, ProSilverScore, HorseScore, DragonScore,
	s.Zero, s.Zero,
	PawnScore, LanceScore, KnightScore, SilverScore, BishopScore, RookScore, GoldScore,
	s.Zero, // King
	ProPawnScore, ProLanceScore, ProKnightScore, ProSilverScore, HorseScore, DragonScore,
}

var CapturedPieceScore = [p.PieceNone]s.Score{
	s.Zero,
	CapturePawnScore, CaptureLanceScore, CaptureKnightScore, CaptureSilverScore, CaptureBishopScore, CaptureRookScore, CaptureGoldScore,
	s.Zero, // King
	CaptureProPawnScore, CaptureProLanceScore, CaptureProKnightScore, CaptureProSilverScore, CaptureHorseScore, CaptureDragonScore,
	s.Zero, s.Zero,
	CapturePawnScore, CaptureLanceScore, CaptureKnightScore, CaptureSilverScore, CaptureBishopScore, CaptureRookScore, CaptureGoldScore,
	s.Zero, // King
	CaptureProPawnScore, CaptureProLanceScore, CaptureProKnightScore, CaptureProSilverScore, CaptureHorseScore, CaptureDragonScore,
}

var PromotePieceScore = [7]s.Score{
	s.Zero,
	PromotePawnScore, PromoteLanceScore, PromoteKnightScore,
	PromoteSilverScore, PromoteBishopScore, PromoteRookScore,
}
