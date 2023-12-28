// Find parameters of stone X (6 vals), such that pos(X) == pos(A)
// at some time, for every stone A.
//
// Optimization/constraint formulation:
// - x,y,z and vx,vy,vz for new stone are variables
// - t values for each stone are variables (not used)
// - constraints are
//       x1 + vx1 * t1 == xX + vxX * t1
//       y1 + vy1 * t1 == yX + vyX * t1
//       z1 + vz1 * t1 == zX + vzX * t1
// - Not really an objective function, except to make
//   sure there is a t for every stone

// Use Centipede constraint solver to solve this

package main

import (
	"context"
	"fmt"
	"math"

	"github.com/gnboorse/centipede"
)

const MaxInt int = 100 // need much higher for input

func part2a() {

	// Create variables x,y,z and vx,vy,vz for the position and
	// velocity of the new stone
	varNames := []centipede.VariableName{"x", "y", "z", "vx", "vy", "vz"}
	vars := centipede.Variables[int]{}
	for _, vname := range varNames {
		v := centipede.NewVariable(vname, centipede.IntRange(1, MaxInt))
		vars = append(vars, v)
	}

	// Create constraint for each stone, such that its position matches
	// that of the new stone, at some time in the future, e.g.,
	//       x1 + vx1 * t1 == xX + vxX * t1
	//       y1 + vy1 * t1 == yX + vyX * t1
	//       z1 + vz1 * t1 == zX + vzX * t1
	constraints := centipede.Constraints[int]{}
	for i, A := range stones {

		// Add a time step variable for this stone
		tVar := centipede.VariableName(fmt.Sprintf("t_%d", i))
		v := centipede.NewVariable(tVar, centipede.IntRange(1, MaxInt))
		vars = append(vars, v)

		// Add the constraint for this stone, that its position
		// is the same as that of the new stone, at some time in
		// the future
		c := centipede.Constraint[int]{Vars: varNames,
			ConstraintFunction: func(variables *centipede.Variables[int]) bool {

				// Return true if any of the variables are empty (why?)
				if variables.Find(tVar).Empty {
					return true
				}
				for _, vn := range varNames {
					if variables.Find(vn).Empty {
						return true
					}
				}

				// Get values of variables
				t := float64(variables.Find(tVar).Value)
				x := float64(variables.Find("x").Value)
				y := float64(variables.Find("y").Value)
				z := float64(variables.Find("z").Value)
				vx := float64(variables.Find("vx").Value)
				vy := float64(variables.Find("vy").Value)
				vz := float64(variables.Find("vz").Value)

				// Position this first stone at this time
				xA := A.x + t*A.vx
				yA := A.y + t*A.vy
				zA := A.z + t*A.vz

				// Position of new stone at this time
				xX := x + t*vx
				yX := y + t*vy
				zX := z + t*vz

				// Check if they match
				return close(xA, xX) && close(yA, yX) && close(zA, zX)
			},
		}

		// Add constraint to the model
		constraints = append(constraints, c)
	}

	// Solve the problem
	fmt.Println(len(vars), "variables, ", len(constraints), "constraints")
	solver := centipede.NewBackTrackingCSPSolver(vars, constraints)
	ctx := context.Background()
	success, _ := solver.Solve(ctx)

	// output results and time elapsed
	if success {
		fmt.Println("Found solution:")
		for _, variable := range solver.State.Vars {
			fmt.Printf("  %v = %v\n", variable.Name, variable.Value)
		}
	} else {
		fmt.Println("Could not find solution")
	}
}

// Determine two numbers are very close
func close(a, b float64) bool {
	return math.Abs(a-b) < .0001
}
