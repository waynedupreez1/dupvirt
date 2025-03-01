/*Package entrypoint
This is the entrypoint of the application

Author: Wayne du Preez
*/
package entrypoint

import (
	"dupvirt/internal/args"
	"dupvirt/internal/logger"
)

func Main(logger logger.ILogger, args *args.Inputs) {
	logger.Info(args.DestinationServer)
}